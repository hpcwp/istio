#!/usr/bin/env python

import requests
import json
import subprocess
import collections
import argparse
import yaml
import sys

POD = collections.namedtuple('Pod', ['name', 'namespace', 'ip', 'labels'])


class XDS(object):

    def __init__(self, url, ns="istio-system", cluster="istio-proxy", headers=None):
        self.url = url
        self.ns = ns
        self.cluster = cluster
        self.headers = headers
        self.cds_info = {}
        self.sds_info = {}

    def key(self, pod):
        role = "sidecar"
        if "ingress" in pod.name:
            role = "ingress"
        elif "egress" in pod.name:
            role = "egress"

        return "{role}~{pod.ip}~{pod.namespace}~{pod.name}.svc.cluster.local".format(
            role=role, pod=pod)

    def query(self, path):
        print path
        try:
            return requests.get(self.url + "/v1" + path, headers=self.headers).json()
        except Exception as ex:
            print ex
            print "Is pilot accessible at ", self.url, " ?"
            sys.exit(-1)

    def lds(self, pod, hydrate=False):
        data = self.query("/listeners/{cluster}/{key}".format(
            cluster=self.cluster, key=self.key(pod)))
        if not hydrate:
            return data

        # call rds

        for l in data['listeners']:
            for f in l['filters']:
                if 'config' not in f:
                    continue
                if 'rds' not in f['config']:
                    continue

                if 'route_config_name' not in f['config']['rds']:
                    continue
                rn = f['config']['rds']['route_config_name']

                # found a route fetch it
                f['config']['route_config'] = self.rds(pod, rn, hydrate)
                f['config']['route_config']['route_config_name'] = rn

        return data

    def rds(self, pod, route="80", hydrate=False):
        data = self.query("/routes/{route}/{cluster}/{key}".format(
            route=route, cluster=self.cluster, key=self.key(pod)))
        if not hydrate:
            return data

        # check if we should hydrate cds
        for vh in data['virtual_hosts']:
            for route in vh['routes']:
                if 'cluster' not in route:
                    continue
                cn = route['cluster']
                route['cluster'] = self.cds(pod, cn, hydrate)
        return data

    def cds(self, pod, cn, hydrate=False):
        pk = self.key(pod)
        if pk not in self.cds_info:
            data = self.query("/clusters/{cluster}/{key}".format(
                cluster=self.cluster, key=self.key(pod)))
            self.cds_info[pk] = {c['name']: c for c in data['clusters']}

            if hydrate:
                for sn, cl in self.cds_info[pk].items():
                    if cl['type'] != "sds":
                        continue
                    cl['endpoints'] = self.sds(cl['service_name'])

        return self.cds_info[pk][cn]

    def sds(self, service_key):
        if service_key not in self.sds_info:
            self.sds_info[service_key] = self.query(
                "/registration/{service_key}".format(service_key=service_key))
        return self.sds_info[service_key]


def pod_info():
    op = subprocess.check_output(
        "kubectl get pod --all-namespaces -o json".split())
    o = json.loads(op)
    return {i['metadata']['name'] + "." + i['metadata']['namespace']:
            POD(i['metadata']['name'], i['metadata']['namespace'],
                i['status']['podIP'], i['metadata']['labels']) for i in o['items']}


def searchpod(pi, searchstr):
    if "," in searchstr:
        si = searchstr.split(',')
        if len(si) != 3:
            print "Use podname,podnamespace,podip format to skip contacting kube api server"
            return None

        return POD(si[0], si[1], si[2], None)

    for pn, pod in pi.items():
        if pn == searchstr:
            return pod
        if searchstr == pod.name:
            return pod
        if searchstr in pod.labels.values():
            return pod

    return None


def main(args):
    pod = searchpod(pod_info(), args.podname)
    if pod is None:
        print "Cound not find pod ", args.podname
        return -1

    xds = XDS(url=args.pilot_url)
    data = xds.lds(pod, True)
    if args.output is None:
        args.output = pod.name + "_xds.yaml"

    op = open(args.output, "wt")
    yaml.safe_dump(data, op, default_flow_style=False, allow_unicode=False)
    print "Wrote ", args.output
    return 0

if __name__ == "__main__":
    parser = argparse.ArgumentParser()

    parser.add_argument("pilot_url",
                        help="Often this is localhost:8080 or 15003 through a port-forward."
                        " \n\nkubectl --namespace=istio-system port-forward $(kubectl --namespace=istio-system get -l istio=pilot pod -o=jsonpath='{.items[0].metadata.name}') 8080:8080"
                        )
    parser.add_argument("podname", help="podname or a label value to search. ingress, mixer, istio-ca all work."
                        " If podname,podnamespace,podip is provided, kubectl is not used to search for the pod")
    parser.add_argument(
        "--output", help="where to write output. default is podname.yaml")
    args = parser.parse_args()
    sys.exit(main(args))

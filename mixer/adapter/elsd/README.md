# ELSd 

This is build following the [walkthrough guide](https://github.com/istio/istio/blob/master/mixer/doc/adapter-development-walkthrough.md)

## Build the mixer server 

```
go build ./...
```

```
go generate $MIXER_REPO/adapter/doc.go
```

## Testing the adapter

Start the Mixer

```shell
pushd $MIXER_REPO && mixs server --configStore2URL=fs://$MIXER_REPO/adapter/elsd/operatorconfig
```

Use the Mixer client to report 

```shell
pushd $MIXER_REPO && go install ./... && mixc report -s="destination.service=svc.cluster.local"
```

Look at the adapter output

```shell
tail $MIXER_REPO/out.txt
```




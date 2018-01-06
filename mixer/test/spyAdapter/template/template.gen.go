// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED.

package template

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/gogo/protobuf/proto"

	"istio.io/api/mixer/v1/config/descriptor"
	adptTmpl "istio.io/api/mixer/v1/template"
	"istio.io/istio/mixer/pkg/adapter"
	"istio.io/istio/mixer/pkg/attribute"
	"istio.io/istio/mixer/pkg/config/proto"
	"istio.io/istio/mixer/pkg/expr"
	"istio.io/istio/mixer/pkg/il/compiled"
	"istio.io/istio/mixer/pkg/template"
	"istio.io/istio/pkg/log"

	"istio.io/istio/mixer/test/spyAdapter/template/apa"

	"istio.io/istio/mixer/test/spyAdapter/template/report"
)

// Add void usages for some imports so that go linter does not complain in case the imports does not get used in the
// below codegen.
var (
	_ net.IP
	_ istio_mixer_v1_config.AttributeManifest
	_ = strings.Reader{}
)

type (
	getFn         func(name string) (value interface{}, found bool)
	namesFn       func() []string
	doneFn        func()
	debugStringFn func() string
	wrapperAttr   struct {
		get         getFn
		names       namesFn
		done        doneFn
		debugString debugStringFn
	}
)

func newWrapperAttrBag(get getFn, names namesFn, done doneFn, debugString debugStringFn) attribute.Bag {
	return &wrapperAttr{
		debugString: debugString,
		done:        done,
		get:         get,
		names:       names,
	}
}

// Get returns an attribute value.
func (w *wrapperAttr) Get(name string) (value interface{}, found bool) {
	return w.get(name)
}

// Names returns the names of all the attributes known to this bag.
func (w *wrapperAttr) Names() []string {
	return w.names()
}

// Done indicates the bag can be reclaimed.
func (w *wrapperAttr) Done() {
	w.done()
}

// DebugString provides a dump of an attribute Bag that avoids affecting the
// calculation of referenced attributes.
func (w *wrapperAttr) DebugString() string {
	return w.debugString()
}

var (
	SupportedTmplInfo = map[string]template.Info{

		sampleapa.TemplateName: {
			Name:               sampleapa.TemplateName,
			Impl:               "sampleapa",
			CtrCfg:             &sampleapa.InstanceParam{},
			Variety:            adptTmpl.TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR,
			BldrInterfaceName:  sampleapa.TemplateName + "." + "HandlerBuilder",
			HndlrInterfaceName: sampleapa.TemplateName + "." + "Handler",
			BuilderSupportsTemplate: func(hndlrBuilder adapter.HandlerBuilder) bool {
				_, ok := hndlrBuilder.(sampleapa.HandlerBuilder)
				return ok
			},
			HandlerSupportsTemplate: func(hndlr adapter.Handler) bool {
				_, ok := hndlr.(sampleapa.Handler)
				return ok
			},
			InferType: func(cp proto.Message, tEvalFn template.TypeEvalFn) (proto.Message, error) {

				var BuildTemplate func(param *sampleapa.InstanceParam,
					path string) (proto.Message, error)

				_ = BuildTemplate

				BuildTemplate = func(param *sampleapa.InstanceParam,
					path string) (proto.Message, error) {

					if param == nil {
						return nil, nil
					}

					var err error = nil

					if param.Int64Primitive == "" {
						return nil, fmt.Errorf("expression for field '%s' cannot be empty", path+"Int64Primitive")
					}
					if t, e := tEvalFn(param.Int64Primitive); e != nil || t != istio_mixer_v1_config_descriptor.INT64 {
						if e != nil {
							return nil, fmt.Errorf("failed to evaluate expression for field '%s': %v", path+"Int64Primitive", e)
						}
						return nil, fmt.Errorf("error type checking for field '%s': Evaluated expression type %v want %v", path+"Int64Primitive", t, istio_mixer_v1_config_descriptor.INT64)
					}

					if param.BoolPrimitive == "" {
						return nil, fmt.Errorf("expression for field '%s' cannot be empty", path+"BoolPrimitive")
					}
					if t, e := tEvalFn(param.BoolPrimitive); e != nil || t != istio_mixer_v1_config_descriptor.BOOL {
						if e != nil {
							return nil, fmt.Errorf("failed to evaluate expression for field '%s': %v", path+"BoolPrimitive", e)
						}
						return nil, fmt.Errorf("error type checking for field '%s': Evaluated expression type %v want %v", path+"BoolPrimitive", t, istio_mixer_v1_config_descriptor.BOOL)
					}

					if param.DoublePrimitive == "" {
						return nil, fmt.Errorf("expression for field '%s' cannot be empty", path+"DoublePrimitive")
					}
					if t, e := tEvalFn(param.DoublePrimitive); e != nil || t != istio_mixer_v1_config_descriptor.DOUBLE {
						if e != nil {
							return nil, fmt.Errorf("failed to evaluate expression for field '%s': %v", path+"DoublePrimitive", e)
						}
						return nil, fmt.Errorf("error type checking for field '%s': Evaluated expression type %v want %v", path+"DoublePrimitive", t, istio_mixer_v1_config_descriptor.DOUBLE)
					}

					if param.StringPrimitive == "" {
						return nil, fmt.Errorf("expression for field '%s' cannot be empty", path+"StringPrimitive")
					}
					if t, e := tEvalFn(param.StringPrimitive); e != nil || t != istio_mixer_v1_config_descriptor.STRING {
						if e != nil {
							return nil, fmt.Errorf("failed to evaluate expression for field '%s': %v", path+"StringPrimitive", e)
						}
						return nil, fmt.Errorf("error type checking for field '%s': Evaluated expression type %v want %v", path+"StringPrimitive", t, istio_mixer_v1_config_descriptor.STRING)
					}

					return nil, err

				}

				instParam := cp.(*sampleapa.InstanceParam)

				const fullOutName = "sampleapa.output."
				for attr, exp := range instParam.AttributeBindings {
					expr := strings.Replace(exp, "$out.", fullOutName, -1)
					t1, err := tEvalFn(expr)
					if err != nil {
						return nil, fmt.Errorf("error evaluating AttributeBinding expression '%s' for attribute '%s': %v", expr, attr, err)
					}
					t2, err := tEvalFn(attr)
					if err != nil {
						return nil, fmt.Errorf("error evaluating AttributeBinding expression for attribute key '%s': %v", attr, err)
					}
					if t1 != t2 {
						return nil, fmt.Errorf(
							"error evaluating AttributeBinding: type '%v' for attribute '%s' does not match type '%s' for expression '%s'",
							t2, attr, t1, expr)
					}
				}

				return BuildTemplate(instParam, "")
			},

			AttributeManifests: []*istio_mixer_v1_config.AttributeManifest{
				{
					Attributes: map[string]*istio_mixer_v1_config.AttributeManifest_AttributeInfo{

						"sampleapa.output.int64Primitive": {
							ValueType: istio_mixer_v1_config_descriptor.INT64,
						},

						"sampleapa.output.boolPrimitive": {
							ValueType: istio_mixer_v1_config_descriptor.BOOL,
						},

						"sampleapa.output.doublePrimitive": {
							ValueType: istio_mixer_v1_config_descriptor.DOUBLE,
						},

						"sampleapa.output.stringPrimitive": {
							ValueType: istio_mixer_v1_config_descriptor.STRING,
						},

						"sampleapa.output.stringMap": {
							ValueType: istio_mixer_v1_config_descriptor.STRING_MAP,
						},
					},
				},
			},

			ProcessGenAttrs: func(ctx context.Context, instName string, inst proto.Message, attrs attribute.Bag,
				mapper expr.Evaluator, handler adapter.Handler) (*attribute.MutableBag, error) {

				var BuildTemplate func(instName string,
					param *sampleapa.InstanceParam, path string) (
					*sampleapa.Instance, error)
				_ = BuildTemplate

				BuildTemplate = func(instName string,
					param *sampleapa.InstanceParam, path string) (
					*sampleapa.Instance, error) {
					if param == nil {
						return nil, nil
					}
					var err error
					_ = err

					Int64Primitive, err := mapper.Eval(param.Int64Primitive, attrs)

					if err != nil {
						msg := fmt.Sprintf("failed to evaluate field '%s' for instance '%s': %v", path+"Int64Primitive", instName, err)
						log.Error(msg)
						return nil, errors.New(msg)
					}

					BoolPrimitive, err := mapper.Eval(param.BoolPrimitive, attrs)

					if err != nil {
						msg := fmt.Sprintf("failed to evaluate field '%s' for instance '%s': %v", path+"BoolPrimitive", instName, err)
						log.Error(msg)
						return nil, errors.New(msg)
					}

					DoublePrimitive, err := mapper.Eval(param.DoublePrimitive, attrs)

					if err != nil {
						msg := fmt.Sprintf("failed to evaluate field '%s' for instance '%s': %v", path+"DoublePrimitive", instName, err)
						log.Error(msg)
						return nil, errors.New(msg)
					}

					StringPrimitive, err := mapper.Eval(param.StringPrimitive, attrs)

					if err != nil {
						msg := fmt.Sprintf("failed to evaluate field '%s' for instance '%s': %v", path+"StringPrimitive", instName, err)
						log.Error(msg)
						return nil, errors.New(msg)
					}

					_ = param
					return &sampleapa.Instance{

						Name: instName,

						Int64Primitive: Int64Primitive.(int64),

						BoolPrimitive: BoolPrimitive.(bool),

						DoublePrimitive: DoublePrimitive.(float64),

						StringPrimitive: StringPrimitive.(string),
					}, nil
				}

				instParam := inst.(*sampleapa.InstanceParam)
				instance, err := BuildTemplate(instName, instParam, "")
				if err != nil {
					return nil, err

				}

				out, err := handler.(sampleapa.Handler).GenerateSampleApaAttributes(ctx, instance)
				if err != nil {
					return nil, err
				}
				abag := attrs
				const fullOutName = "sampleapa.output."
				if out == nil {
					log.Debugf("Preprocess adapter returned nil output for instance name '%s'", instName)
				} else {
					abag = newWrapperAttrBag(
						func(name string) (value interface{}, found bool) {
							field := strings.TrimPrefix(name, fullOutName)
							if len(field) != len(name) {
								switch field {

								case "int64Primitive":

									return out.Int64Primitive, true

								case "boolPrimitive":

									return out.BoolPrimitive, true

								case "doublePrimitive":

									return out.DoublePrimitive, true

								case "stringPrimitive":

									return out.StringPrimitive, true

								case "stringMap":

									return out.StringMap, true

								default:
									return nil, false
								}

							}
							return attrs.Get(name)
						},
						func() []string { return attrs.Names() },
						func() { attrs.Done() },
						func() string { return attrs.DebugString() },
					)
				}
				resultBag := attribute.GetMutableBag(nil)
				for attrName, outExpr := range instParam.AttributeBindings {
					ex := strings.Replace(outExpr, "$out.", fullOutName, -1)
					val, err := mapper.Eval(ex, abag)
					if err != nil {
						return nil, err
					}
					switch v := val.(type) {
					case net.IP:
						// conversion to []byte necessary based on current IP_ADDRESS handling within Mixer
						// TODO: remove
						if v4 := v.To4(); v4 != nil {
							resultBag.Set(attrName, []byte(v4))
							continue
						}
						resultBag.Set(attrName, []byte(v.To16()))
					default:
						resultBag.Set(attrName, val)
					}
				}
				return resultBag, nil

			},

			/* runtime2 bindings */

			// DispathGenAttrs dispatches the instance to the attribute producing handler.
			DispatchGenAttrs: func(ctx context.Context, handler adapter.Handler, inst interface{}, attrs attribute.Bag,
				mapper template.OutputMapperFn) (*attribute.MutableBag, error) {

				// Convert the instance from the generic interface{}, to their specialized type.
				instance := inst.(*sampleapa.Instance)

				// Invoke the handler.
				out, err := handler.(sampleapa.Handler).GenerateSampleApaAttributes(ctx, instance)
				if err != nil {
					return nil, err
				}

				// Construct a wrapper bag around the returned output message and pass it to the output mapper
				// to map $out values back to the destination attributes in the ambient context.
				const fullOutName = "sampleapa.output."
				outBag := newWrapperAttrBag(
					func(name string) (value interface{}, found bool) {
						field := strings.TrimPrefix(name, fullOutName)
						if len(field) != len(name) {
							switch field {

							case "int64Primitive":

								return out.Int64Primitive, true

							case "boolPrimitive":

								return out.BoolPrimitive, true

							case "doublePrimitive":

								return out.DoublePrimitive, true

							case "stringPrimitive":

								return out.StringPrimitive, true

							case "stringMap":

								return out.StringMap, true

							default:
								return nil, false
							}
						}
						return attrs.Get(name)
					},
					func() []string { return attrs.Names() },
					func() { attrs.Done() },
					func() string { return attrs.DebugString() },
				)

				// Mapper will map back $out values in the outBag into ambient attribute names, and return
				// a bag with these additional attributes.
				return mapper(outBag)
			},

			// CreateInstanceBuilder creates a new template.InstanceBuilderFN based on the supplied instance parameters. It uses
			// the expression builder to create a new instance of a builder struct for the instance type. Created
			// InstanceBuilderFn closes over this struct. When InstanceBuilderFn is called it, in turn, calls into
			// the builder with an attribute bag.
			//
			// See template.CreateInstanceBuilderFn for more details.
			CreateInstanceBuilder: func(instanceName string, param proto.Message, expb *compiled.ExpressionBuilder) (template.InstanceBuilderFn, error) {

				// If the parameter is nil. Simply return nil. The builder, then, will also return nil.
				if param == nil {
					return func(attr attribute.Bag) (interface{}, error) {
						return nil, nil
					}, nil
				}

				// Instantiate a new builder for the instance.
				builder, errp := newBuilder_sampleapa_Template(expb, param.(*sampleapa.InstanceParam))
				if !errp.IsNil() {
					return nil, errp.AsCompilationError(instanceName)
				}

				return func(attr attribute.Bag) (interface{}, error) {
					// Use the instantiated builder (that this fn closes over) to construct an instance.
					e, errp := builder.build(attr)
					if !errp.IsNil() {
						err := errp.AsEvaluationError(instanceName)
						log.Error(err.Error())
						return nil, err
					}

					e.Name = instanceName
					return e, nil
				}, nil
			},

			// CreateOutputExpressions creates a set of compiled expressions based on the supplied instance parameters.
			//
			// See template.CreateOutputExpressionsFn for more details.
			CreateOutputExpressions: func(
				instanceParam proto.Message,
				finder expr.AttributeDescriptorFinder,
				expb *compiled.ExpressionBuilder) (map[string]compiled.Expression, error) {
				var err error
				var expType istio_mixer_v1_config_descriptor.ValueType

				// Convert the generic instanceParam to its specialized type.
				param := instanceParam.(*sampleapa.InstanceParam)

				// Create a mapping of expressions back to the attribute names.
				expressions := make(map[string]compiled.Expression, len(param.AttributeBindings))

				const fullOutName = "sampleapa.output."
				for attrName, outExpr := range param.AttributeBindings {
					attrInfo := finder.GetAttribute(attrName)
					if attrInfo == nil {
						log.Warnf("attribute not found when mapping outputs: attr='%s', expr='%s'", attrName, outExpr)
						continue
					}

					ex := strings.Replace(outExpr, "$out.", fullOutName, -1)

					if expressions[attrName], expType, err = expb.Compile(ex); err != nil {
						return nil, err
					}

					if attrInfo.ValueType != expType {
						log.Warnf("attribute type mismatch: attr='%s', attrType='%v', expr='%s', exprType='%v'", attrName, attrInfo.ValueType, outExpr, expType)
						continue
					}
				}

				return expressions, nil
			},
		},

		samplereport.TemplateName: {
			Name:               samplereport.TemplateName,
			Impl:               "samplereport",
			CtrCfg:             &samplereport.InstanceParam{},
			Variety:            adptTmpl.TEMPLATE_VARIETY_REPORT,
			BldrInterfaceName:  samplereport.TemplateName + "." + "HandlerBuilder",
			HndlrInterfaceName: samplereport.TemplateName + "." + "Handler",
			BuilderSupportsTemplate: func(hndlrBuilder adapter.HandlerBuilder) bool {
				_, ok := hndlrBuilder.(samplereport.HandlerBuilder)
				return ok
			},
			HandlerSupportsTemplate: func(hndlr adapter.Handler) bool {
				_, ok := hndlr.(samplereport.Handler)
				return ok
			},
			InferType: func(cp proto.Message, tEvalFn template.TypeEvalFn) (proto.Message, error) {

				var BuildTemplate func(param *samplereport.InstanceParam,
					path string) (*samplereport.Type, error)

				_ = BuildTemplate

				BuildTemplate = func(param *samplereport.InstanceParam,
					path string) (*samplereport.Type, error) {

					if param == nil {
						return nil, nil
					}

					infrdType := &samplereport.Type{}

					var err error = nil

					if param.Value == "" {
						return nil, fmt.Errorf("expression for field '%s' cannot be empty", path+"Value")
					}
					if infrdType.Value, err = tEvalFn(param.Value); err != nil {
						return nil, fmt.Errorf("failed to evaluate expression for field '%s'; %v", path+"Value", err)
					}

					infrdType.Dimensions = make(map[string]istio_mixer_v1_config_descriptor.ValueType, len(param.Dimensions))

					for k, v := range param.Dimensions {

						if infrdType.Dimensions[k], err = tEvalFn(v); err != nil {

							return nil, fmt.Errorf("failed to evaluate expression for field '%s'; %v", path+"Dimensions", err)
						}
					}

					return infrdType, err

				}

				instParam := cp.(*samplereport.InstanceParam)

				return BuildTemplate(instParam, "")
			},

			SetType: func(types map[string]proto.Message, builder adapter.HandlerBuilder) {
				// Mixer framework should have ensured the type safety.
				castedBuilder := builder.(samplereport.HandlerBuilder)
				castedTypes := make(map[string]*samplereport.Type, len(types))
				for k, v := range types {
					// Mixer framework should have ensured the type safety.
					v1 := v.(*samplereport.Type)
					castedTypes[k] = v1
				}
				castedBuilder.SetSampleReportTypes(castedTypes)
			},

			ProcessReport: func(ctx context.Context, insts map[string]proto.Message, attrs attribute.Bag, mapper expr.Evaluator, handler adapter.Handler) error {

				var BuildTemplate func(instName string,
					param *samplereport.InstanceParam, path string) (
					*samplereport.Instance, error)
				_ = BuildTemplate

				BuildTemplate = func(instName string,
					param *samplereport.InstanceParam, path string) (
					*samplereport.Instance, error) {
					if param == nil {
						return nil, nil
					}
					var err error
					_ = err

					Value, err := mapper.Eval(param.Value, attrs)

					if err != nil {
						msg := fmt.Sprintf("failed to evaluate field '%s' for instance '%s': %v", path+"Value", instName, err)
						log.Error(msg)
						return nil, errors.New(msg)
					}

					Dimensions, err := template.EvalAll(param.Dimensions, attrs, mapper)

					if err != nil {
						msg := fmt.Sprintf("failed to evaluate field '%s' for instance '%s': %v", path+"Dimensions", instName, err)
						log.Error(msg)
						return nil, errors.New(msg)
					}

					_ = param
					return &samplereport.Instance{

						Name: instName,

						Value: Value,

						Dimensions: Dimensions,
					}, nil
				}

				var instances []*samplereport.Instance
				for instName, inst := range insts {
					instance, err := BuildTemplate(instName, inst.(*samplereport.InstanceParam), "")
					if err != nil {
						return err
					}
					instances = append(instances, instance)
				}

				if err := handler.(samplereport.Handler).HandleSampleReport(ctx, instances); err != nil {
					return fmt.Errorf("failed to report all values: %v", err)
				}
				return nil
			},

			/* runtime2 bindings */

			// DispatchReport dispatches the instances to the handler.
			DispatchReport: func(ctx context.Context, handler adapter.Handler, inst []interface{}) error {

				// Convert the instances from the generic []interface{}, to their specialized type.
				instances := make([]*samplereport.Instance, len(inst))
				for i, instance := range inst {
					instances[i] = instance.(*samplereport.Instance)
				}

				// Invoke the handler.
				if err := handler.(samplereport.Handler).HandleSampleReport(ctx, instances); err != nil {
					return fmt.Errorf("failed to report all values: %v", err)
				}
				return nil
			},

			// CreateInstanceBuilder creates a new template.InstanceBuilderFN based on the supplied instance parameters. It uses
			// the expression builder to create a new instance of a builder struct for the instance type. Created
			// InstanceBuilderFn closes over this struct. When InstanceBuilderFn is called it, in turn, calls into
			// the builder with an attribute bag.
			//
			// See template.CreateInstanceBuilderFn for more details.
			CreateInstanceBuilder: func(instanceName string, param proto.Message, expb *compiled.ExpressionBuilder) (template.InstanceBuilderFn, error) {

				// If the parameter is nil. Simply return nil. The builder, then, will also return nil.
				if param == nil {
					return func(attr attribute.Bag) (interface{}, error) {
						return nil, nil
					}, nil
				}

				// Instantiate a new builder for the instance.
				builder, errp := newBuilder_samplereport_Template(expb, param.(*samplereport.InstanceParam))
				if !errp.IsNil() {
					return nil, errp.AsCompilationError(instanceName)
				}

				return func(attr attribute.Bag) (interface{}, error) {
					// Use the instantiated builder (that this fn closes over) to construct an instance.
					e, errp := builder.build(attr)
					if !errp.IsNil() {
						err := errp.AsEvaluationError(instanceName)
						log.Error(err.Error())
						return nil, err
					}

					e.Name = instanceName
					return e, nil
				}, nil
			},
		},
	}
)

// Builders for all known message types.

// builder struct for constructing an instance of Template.
type builder_sampleapa_Template struct {

	// builder for field int64Primitive: int64.

	bldInt64Primitive compiled.Expression

	// builder for field boolPrimitive: bool.

	bldBoolPrimitive compiled.Expression

	// builder for field doublePrimitive: float64.

	bldDoublePrimitive compiled.Expression

	// builder for field stringPrimitive: string.

	bldStringPrimitive compiled.Expression
} // builder_sampleapa_Template

// Instantiates and returns a new builder for Template, based on the provided instance parameter.
func newBuilder_sampleapa_Template(
	expb *compiled.ExpressionBuilder,
	param *sampleapa.InstanceParam) (*builder_sampleapa_Template, template.ErrorPath) {

	// If the parameter is nil. Simply return nil. The builder, then, will also return nil.
	if param == nil {
		return nil, template.ErrorPath{}
	}

	b := &builder_sampleapa_Template{}

	var exp compiled.Expression
	_ = exp
	var err error
	_ = err
	var errp template.ErrorPath
	_ = errp
	var expType istio_mixer_v1_config_descriptor.ValueType
	_ = expType

	b.bldInt64Primitive, expType, err = expb.Compile(param.Int64Primitive)
	if err != nil {
		return nil, template.NewErrorPath("Int64Primitive", err)
	}

	if expType != istio_mixer_v1_config_descriptor.INT64 {
		err = fmt.Errorf("instance field type mismatch: expected='%v', actual='%v', expression='%s'", istio_mixer_v1_config_descriptor.INT64, expType, param.Int64Primitive)
		return nil, template.NewErrorPath("Int64Primitive", err)
	}

	b.bldBoolPrimitive, expType, err = expb.Compile(param.BoolPrimitive)
	if err != nil {
		return nil, template.NewErrorPath("BoolPrimitive", err)
	}

	if expType != istio_mixer_v1_config_descriptor.BOOL {
		err = fmt.Errorf("instance field type mismatch: expected='%v', actual='%v', expression='%s'", istio_mixer_v1_config_descriptor.BOOL, expType, param.BoolPrimitive)
		return nil, template.NewErrorPath("BoolPrimitive", err)
	}

	b.bldDoublePrimitive, expType, err = expb.Compile(param.DoublePrimitive)
	if err != nil {
		return nil, template.NewErrorPath("DoublePrimitive", err)
	}

	if expType != istio_mixer_v1_config_descriptor.DOUBLE {
		err = fmt.Errorf("instance field type mismatch: expected='%v', actual='%v', expression='%s'", istio_mixer_v1_config_descriptor.DOUBLE, expType, param.DoublePrimitive)
		return nil, template.NewErrorPath("DoublePrimitive", err)
	}

	b.bldStringPrimitive, expType, err = expb.Compile(param.StringPrimitive)
	if err != nil {
		return nil, template.NewErrorPath("StringPrimitive", err)
	}

	if expType != istio_mixer_v1_config_descriptor.STRING {
		err = fmt.Errorf("instance field type mismatch: expected='%v', actual='%v', expression='%s'", istio_mixer_v1_config_descriptor.STRING, expType, param.StringPrimitive)
		return nil, template.NewErrorPath("StringPrimitive", err)
	}

	return b, template.ErrorPath{}
}

// build and return the instance, given a set of attributes.
func (b *builder_sampleapa_Template) build(
	attrs attribute.Bag) (*sampleapa.Instance, template.ErrorPath) {

	if b == nil {
		return nil, template.ErrorPath{}
	}

	var err error
	_ = err
	var errp template.ErrorPath
	_ = errp
	var vBool bool
	_ = vBool
	var vInt int64
	_ = vInt
	var vString string
	_ = vString
	var vDouble float64
	_ = vDouble
	var vIface interface{}
	_ = vIface

	r := &sampleapa.Instance{}

	vInt, err = b.bldInt64Primitive.EvaluateInteger(attrs)
	if err != nil {
		return nil, template.NewErrorPath("Int64Primitive", err)
	}
	r.Int64Primitive = vInt

	vBool, err = b.bldBoolPrimitive.EvaluateBoolean(attrs)
	if err != nil {
		return nil, template.NewErrorPath("BoolPrimitive", err)
	}
	r.BoolPrimitive = vBool

	vDouble, err = b.bldDoublePrimitive.EvaluateDouble(attrs)
	if err != nil {
		return nil, template.NewErrorPath("DoublePrimitive", err)
	}
	r.DoublePrimitive = vDouble

	vString, err = b.bldStringPrimitive.EvaluateString(attrs)
	if err != nil {
		return nil, template.NewErrorPath("StringPrimitive", err)
	}
	r.StringPrimitive = vString

	return r, template.ErrorPath{}
}

// builder struct for constructing an instance of Template.
type builder_samplereport_Template struct {

	// builder for field value: istio_mixer_v1_config_descriptor.ValueType.

	bldValue compiled.Expression

	// builder for field dimensions: map[string]istio_mixer_v1_config_descriptor.ValueType.

	bldDimensions map[string]compiled.Expression
} // builder_samplereport_Template

// Instantiates and returns a new builder for Template, based on the provided instance parameter.
func newBuilder_samplereport_Template(
	expb *compiled.ExpressionBuilder,
	param *samplereport.InstanceParam) (*builder_samplereport_Template, template.ErrorPath) {

	// If the parameter is nil. Simply return nil. The builder, then, will also return nil.
	if param == nil {
		return nil, template.ErrorPath{}
	}

	b := &builder_samplereport_Template{}

	var exp compiled.Expression
	_ = exp
	var err error
	_ = err
	var errp template.ErrorPath
	_ = errp
	var expType istio_mixer_v1_config_descriptor.ValueType
	_ = expType

	b.bldValue, expType, err = expb.Compile(param.Value)
	if err != nil {
		return nil, template.NewErrorPath("Value", err)
	}

	b.bldDimensions = make(map[string]compiled.Expression, len(param.Dimensions))
	for k, v := range param.Dimensions {
		var exp compiled.Expression
		if exp, expType, err = expb.Compile(v); err != nil {
			return nil, template.NewErrorPath("Dimensions["+k+"]", err)
		}

		b.bldDimensions[k] = exp
	}

	return b, template.ErrorPath{}
}

// build and return the instance, given a set of attributes.
func (b *builder_samplereport_Template) build(
	attrs attribute.Bag) (*samplereport.Instance, template.ErrorPath) {

	if b == nil {
		return nil, template.ErrorPath{}
	}

	var err error
	_ = err
	var errp template.ErrorPath
	_ = errp
	var vBool bool
	_ = vBool
	var vInt int64
	_ = vInt
	var vString string
	_ = vString
	var vDouble float64
	_ = vDouble
	var vIface interface{}
	_ = vIface

	r := &samplereport.Instance{}

	if vIface, err = b.bldValue.Evaluate(attrs); err != nil {
		return nil, template.NewErrorPath("Value", err)
	}

	r.Value = vIface

	r.Dimensions = make(map[string]interface{}, len(b.bldDimensions))

	for k, v := range b.bldDimensions {

		if vIface, err = v.Evaluate(attrs); err != nil {
			return nil, template.NewErrorPath("Dimensions["+k+"]", err)
		}

		r.Dimensions[k] = vIface

	}

	return r, template.ErrorPath{}
}

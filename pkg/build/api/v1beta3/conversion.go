package v1beta3

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/conversion"

	newer "github.com/openshift/origin/pkg/build/api"
)

func init() {
	api.Scheme.AddConversionFuncs(
		func(in *newer.Build, out *Build, s conversion.Scope) error {
			if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Parameters, &out.Spec, 0); err != nil {
				return err
			}
			if err := s.Convert(in, &out.Status, 0); err != nil {
				return err
			}
			return s.Convert(&in.Status, &out.Status.Phase, 0)
		},
		func(in *Build, out *newer.Build, s conversion.Scope) error {
			if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Status, out, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Spec, &out.Parameters, 0); err != nil {
				return err
			}
			return s.Convert(&in.Status.Phase, &out.Status, 0)
		},

		func(in *newer.Build, out *BuildStatus, s conversion.Scope) error {
			out.Cancelled = in.Cancelled
			out.CompletionTimestamp = in.CompletionTimestamp
			if err := s.Convert(&in.Config, &out.Config, 0); err != nil {
				return err
			}
			out.Duration = in.Duration
			out.Message = in.Message
			out.StartTimestamp = in.StartTimestamp
			return nil
		},
		func(in *BuildStatus, out *newer.Build, s conversion.Scope) error {
			out.Cancelled = in.Cancelled
			out.CompletionTimestamp = in.CompletionTimestamp
			if err := s.Convert(&in.Config, &out.Config, 0); err != nil {
				return err
			}
			out.Duration = in.Duration
			out.Message = in.Message
			out.StartTimestamp = in.StartTimestamp
			return nil
		},

		func(in *newer.BuildStatus, out *BuildPhase, s conversion.Scope) error {
			str := string(*in)
			*out = BuildPhase(str)
			return nil
		},
		func(in *BuildPhase, out *newer.BuildStatus, s conversion.Scope) error {
			str := string(*in)
			*out = newer.BuildStatus(str)
			return nil
		},

		func(in *newer.BuildConfig, out *BuildConfig, s conversion.Scope) error {
			if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
				return err
			}
			if err := s.Convert(&in.Parameters, &out.Spec, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Triggers, &out.Spec.Triggers, 0); err != nil {
				return err
			}
			out.Status.LastVersion = in.LastVersion
			return nil
		},
		func(in *BuildConfig, out *newer.BuildConfig, s conversion.Scope) error {
			if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
				return err
			}
			if err := s.Convert(&in.Spec, &out.Parameters, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Spec.Triggers, &out.Triggers, 0); err != nil {
				return err
			}
			out.LastVersion = in.Status.LastVersion
			return nil
		},

		func(in *newer.BuildParameters, out *BuildSpec, s conversion.Scope) error {
			err := s.DefaultConvert(&in.Strategy, &out.Strategy, conversion.IgnoreMissingFields)
			if err != nil {
				return err
			}
			if err := s.Convert(&in.Source, &out.Source, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Output, &out.Output, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Revision, &out.Revision, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Resources, &out.Resources, 0); err != nil {
				return err
			}
			return nil
		},
		func(in *BuildSpec, out *newer.BuildParameters, s conversion.Scope) error {
			err := s.DefaultConvert(&in.Strategy, &out.Strategy, conversion.IgnoreMissingFields)
			if err != nil {
				return err
			}
			if err := s.Convert(&in.Source, &out.Source, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Output, &out.Output, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Revision, &out.Revision, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Resources, &out.Resources, 0); err != nil {
				return err
			}
			return nil
		},

		func(in *newer.BuildParameters, out *BuildConfigSpec, s conversion.Scope) error {
			err := s.DefaultConvert(&in.Strategy, &out.Strategy, conversion.IgnoreMissingFields)
			if err != nil {
				return err
			}
			if err := s.Convert(&in.Source, &out.Source, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Output, &out.Output, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Revision, &out.Revision, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Resources, &out.Resources, 0); err != nil {
				return err
			}
			return nil
		},
		func(in *BuildConfigSpec, out *newer.BuildParameters, s conversion.Scope) error {
			err := s.DefaultConvert(&in.Strategy, &out.Strategy, conversion.IgnoreMissingFields)
			if err != nil {
				return err
			}
			if err := s.Convert(&in.Source, &out.Source, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Output, &out.Output, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Revision, &out.Revision, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Resources, &out.Resources, 0); err != nil {
				return err
			}
			return nil
		},

		func(in *newer.STIBuildStrategy, out *STIBuildStrategy, s conversion.Scope) error {
			return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
		},
		func(in *STIBuildStrategy, out *newer.STIBuildStrategy, s conversion.Scope) error {
			return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
		},
		func(in *newer.DockerBuildStrategy, out *DockerBuildStrategy, s conversion.Scope) error {
			return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
		},
		func(in *DockerBuildStrategy, out *newer.DockerBuildStrategy, s conversion.Scope) error {
			return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
		},
		func(in *newer.CustomBuildStrategy, out *CustomBuildStrategy, s conversion.Scope) error {
			return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
		},
		func(in *CustomBuildStrategy, out *newer.CustomBuildStrategy, s conversion.Scope) error {
			return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
		},
		func(in *newer.BuildOutput, out *BuildOutput, s conversion.Scope) error {
			return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
		},
		func(in *BuildOutput, out *newer.BuildOutput, s conversion.Scope) error {
			return s.DefaultConvert(in, out, conversion.IgnoreMissingFields)
		},
		func(in *newer.ImageChangeTrigger, out *ImageChangeTrigger, s conversion.Scope) error {
			out.LastTriggeredImageID = in.LastTriggeredImageID
			return nil
		},
		func(in *ImageChangeTrigger, out *newer.ImageChangeTrigger, s conversion.Scope) error {
			out.LastTriggeredImageID = in.LastTriggeredImageID
			return nil
		})
}

package tasks

import "reflect"

type OrchestratableTaskInputArgs interface {
    Validate() error
}

type OrchestratableTaskOutput interface{}

type OrchestratableTask interface {
    Run(input OrchestratableTaskInputArgs) (OrchestratableTaskOutput, error)
    InputType() reflect.Type
    OutputType() reflect.Type
}

type BaseOrchestratableTask struct {
    inputType  reflect.Type
    outputType reflect.Type
}

func NewBaseOrchestratableTask(input OrchestratableTaskInputArgs, output OrchestratableTaskOutput) BaseOrchestratableTask {
    return BaseOrchestratableTask{
        inputType:  reflect.TypeOf(input),
        outputType: reflect.TypeOf(output),
    }
}

func (b *BaseOrchestratableTask) InputType() reflect.Type {
    return b.inputType
}

func (b *BaseOrchestratableTask) OutputType() reflect.Type {
    return b.outputType
}

package znet

import "zinx/src/ziface"

// 实现其它 Router 时, 可以先嵌入这个基类, 然后根据需要对这个基类的方法进行重写
// 基于这种方式, 对于具体的 Router, 如果不需要 PostHandle / PreHandle 则可以不用实现
// 因为它们会默认使用基类的 PreHandle / PostHandle
type BaseRouter struct{}

func (br *BaseRouter) PreHandle(req ziface.IRequest)  {}
func (br *BaseRouter) Handle(req ziface.IRequest)     {}
func (br *BaseRouter) PostHandle(req ziface.IRequest) {}

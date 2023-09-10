# MBTI人格测试网站

---

## 📃 系统需求

- 构建一个MBTI测试系统记录每一个学院、专业和班级的学生测试结果；
- 建立组织结构管理；
- 建立MBTI题库系统；
- 建立不同时间的学生测试结果管理系统；
- 个人和班级测试结果可视化。



---



## 📝 系统介绍

分为四个子系统：用户系统、题库系统、测试数据系统、学校管理系统。

用户系统：

>实现登录、注册、获取用户信息功能。

题库系统：

>实现新增、编辑、删除题目，获取题目信息功能。

测试数据系统：

>实现数据录入（包括每一条测试的选项详细信息），获取测试数据功能。

学校管理系统：

>主要模拟学校提供的学生信息接口，方便进行学院、专业、班级测试结果可视化。



---



## ⚙ 技术栈

视图层：Vue3 + VueCLI

​	一套构建用户界面的渐进式框架，只关注视图层， 采用自底向上增量开发的设计。

控制层：Go-Zero微服务框架

- 强大的工具支持，尽可能少的代码编写
- 极简的接口
- 完全兼容net/http
- 支持中间件，方便扩展
- 高性能
- 面向故障编程，弹性设计
- 内建服务发现、负载均衡
- 内建限流、熔断、降载，且自动触发，自动恢复
- API参数自动校验
- 超时级联控制
- 自动缓存控制
- 链路跟踪、统计报警等
- 高并发支撑，稳定保障了流量洪峰

持久化：Redis + PostgreSQL

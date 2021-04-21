<p align="center" style="margin: 20px 0 40px 0;">
  <img height="100" src="https://zoom-ci.github.io/docs/assets/img/logo_black.png" />
</p>
<h3 align="center">Zoom - A CI/CD Service.</h3>

<p align="center">Zoom是一款持续集成、部署服务工具，它具有简单、高效、易用等特点，可以提高团队的工作效率。</p>


## 特性

- Go语言开发，使用简单、运行高效
- 权限模型灵活自由
- 支持远程和本地两种部署方式（可实现客户端打包）
- 支持自定义构建
- 支持Git仓库
- 支持分支、Tag上线
- 部署Hook支持，可扩展性强
- 完善的上线工作流
- 邮件通知机制

## 安装说明

1、下载[最新版本release包](https://github.com/zoom-ci/zoom-ci/releases),并将其拷贝到任意目录（比如：~/zoom_workspace）并执行;

```shell
$ ./zoom-v1.0.0-darwin-amd64   # 这里以mac 64位版为例 

 _____________________________
       ___                    
         /                    
 -------/-----__----__---_--_-
       /    /   ) /   ) / /  )
 ____(_____(___/_(___/_/_/__/_
     / Zoom,a CI/CD service.  
 (_ /                         


Service:              zoom
Version:              v1.0.0
Config Loaded:        ./zoom.ini
Log:                  stdout
Mail Enable:          0
HTTP Service:         :7002
Start Running...
```

2、打开浏览器，访问 `http://localhost:7002` (出现下图界面)，配置数据库，安装完成。
<p style="margin: 20px 0 40px 0;">
  <img height="500"  src="https://zoom-ci.github.io/docs/assets/img/zoom-install.png" />
</p>

初始账号：
```
用户名: admin
密码: 123456
```

## 文档

#### [https://zoom-ci.github.io/docs/](https://zoom-ci.github.io/docs/)

## TODO

- 安装流程简化（已完成）
- 支持项目复制（已完成）
- 支持远程、本地模式（已完成） 
- 支持定时任务
- 支持WebHook调用

## QQ群
<p style="margin: 20px 0 40px 0;">
  <img height="200" src="https://zoom-ci.github.io/docs/assets/img/qq.png" />
</p>


## LICENSE

本项目采用 MIT 开源授权许可证，完整的授权说明已放置在 LICENSE 文件中。
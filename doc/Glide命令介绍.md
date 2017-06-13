#glide create (别名 init)

初始化新工作区。除此之外，这会创建一个`glide.yaml`文件，同时试图猜测`package`和版本。例如，如果你的项目使用`Godep`，它将使用`Godep`指定的版本。`Glide`足够智能去扫描您的代码库，检测正在使用的`package`，无论有没有指定其他的包管理器。
```
$ glide create
[INFO]  Generating a YAML configuration file and guessing the dependencies
[INFO]  Attempting to import from other package managers (use --skip-import to skip)
[INFO]  Scanning code to look for dependencies
[INFO]  --> Found reference to github.com/Masterminds/semver
[INFO]  --> Found reference to github.com/Masterminds/vcs
[INFO]  --> Found reference to github.com/codegangsta/cli
[INFO]  --> Found reference to gopkg.in/yaml.v2
[INFO]  Writing configuration file (glide.yaml)
[INFO]  Would you like Glide to help you find ways to improve your glide.yaml configuration?
[INFO]  If you want to revisit this step you can use the config-wizard command at any time.
[INFO]  Yes (Y) or No (N)?
n
[INFO]  You can now edit the glide.yaml file. Consider:
[INFO]  --> Using versions and ranges. See https://glide.sh/docs/versions/
[INFO]  --> Adding additional metadata. See https://glide.sh/docs/glide.yaml/
[INFO]  --> Running the config-wizard command to improve the versions in your configuration
```
这里提到的配置向导可以在这里运行或者以后手动运行。此向导可帮助您找出可用于依赖项的版本和范围。

#glide config-wizard

这将运行一个向导，扫描依赖关系并检索其上的信息，以提供可以交互选择的建议。例如，它可以发现依赖关系是否使用语义版本，并帮助您选择要使用的版本范围。

#glide get [package name]

你可以通过`glide get`下载一个或多个包到你的`vendor`目录，并会自动加入到`glide.yml`文件中。
```
$ glide get github.com/Masterminds/cookoo
```
当使用`glide get`时，它将内省所列出的软件包来解决它的依赖性，包括使用`Godep`，`GPM`，`Gom`和`GB`配置文件。

`glide get`命令可以使用包名称传递一个版本或范围。 例如：
```
$ glide get github.com/Masterminds/cookoo#^1.2.3
```
版本通过锚（`＃`）与包名称分隔开。如果未指定版本或范围，并且依赖关系使用语义版本 `Glide` 将提示您询问是否要使用它们。

#glide update (别名 up)

下载或更新`glide.yml`文件中列出的所有库，并将它们放在`vendor`目录中。它还将递归遍历依赖包以获取任何所需的配置并在任何配置中读取。
```
$ glide up
```
这将递归在寻找由`Glide`，`Godep`，`gb`，`gom`和`GPM`管理的其他项目的包。当找到这些包时，将根据需要安装这些包。

将创建或更新`glide.lock`文件，并将依赖关系固定到特定版本。例如，如果在`glide.yaml`文件中将版本指定为范围（例如，`^ 1.2.3`），它将被设置为`glide.lock`文件中的特定提交标识。这允许可重复安装（请参阅`glide install`）。

要从已提取的包中删除任何嵌套的`vendor/`目录，请参见`-v`标志。

#glide install

当你想从`glide.lock`文件安装特定的版本使用`glide install`。
```
$ glide install
```
这将读取`glide.lock`文件，警告你如果它没有绑定到`glide.yaml`文件，并安装`commit id`特定的版本。

当`glide.lock`文件不绑定到`glide.yaml`文件时，如有更改，它将提供警告。运行`glide up`将在更新依赖关系树时重新创建`glide.lock`文件。

如果没有`glide.lock`文件存在`glide install`将执行`update`并生成 `lock` 文件。

要从已提取的包中删除任何嵌套的`vendor/`目录，请参见`-v`标志。


#glide novendor (别名 nv)

当你运行`go test ./...`这样的命令时，它会遍历所有的子目录，包括`vendor`目录。当你测试你的应用程序时，你可能想测试你的应用程序文件，而不需要运行所有依赖项及其依赖关系的测试。这是`novendor`命令进来的地方。它列出除了`vendor`以外的所有目录。
```
$ go test $(glide novendor)
```
这将对您的项目的所有目录（`vendor`目录除外）运行`go test`。


#glide name

当你使用 `Glide` 编写脚本时，有时你需要知道你正在使用的包的名称。`glide name`返回`glide.yaml`文件中列出的软件包的名称。


#glide list

`Glide` 的`list`命令显示项目导入的所有包的按字母顺序排列的列表。
```
$ glide list
INSTALLED packages:
    vendor/github.com/Masterminds/cookoo
    vendor/github.com/Masterminds/cookoo/fmt
    vendor/github.com/Masterminds/cookoo/io
    vendor/github.com/Masterminds/cookoo/web
    vendor/github.com/Masterminds/semver
    vendor/github.com/Masterminds/vcs
    vendor/github.com/codegangsta/cli
    vendor/gopkg.in/yaml.v2
```

#glide help

打印 `glide` 帮助
```
$ glide help
```

#glide –version

显示版本信息
```
$ glide --version
glide version 0.12.0
```


#glide mirror

镜像提供了将 `repo` 位置替换为作为原始镜像的另一位置的能力。当您希望拥有连续集成（`CI`）系统的缓存时，或者如果您要在本地位置的依赖项上工作时，这是非常有用的。

镜像存储在`GLIDE_HOME`中的`mirrors.yaml`文件中。

到管理器镜像的三个命令是`list`，`set`和`remove`。

在表单中使用`set`：
```
glide mirror set [original] [replacement]
```
或
```
glide mirror set [original] [replacement] --vcs [type]
```
例如：
```
$ glide mirror set https://github.com/example/foo https://git.example.com/example/foo.git
$ glide mirror set https://github.com/example/foo file:///path/to/local/repo --vcs git
```
请在表单中使用`remove`：
```
glide mirror remove [original]
```
例如：
```
$ glide mirror remove https://github.com/example/foo
```

#参考链接

[Golang包管理工具glide简介](http://www.cnblogs.com/xiwang/p/5870941.html)
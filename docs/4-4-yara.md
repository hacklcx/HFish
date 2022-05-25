#### HFish检测语法介绍

HFish使用的yara监测语法分为三部分，rule，string和condition。样例如下：

```
rule name1
{

strings:
$str1 = "1"
$str12 = "2"

condition:
any of them
}
```

该样例代表的是，我建立了一个叫做"name1"的检测规则，如果我检测的数据中出现1或者2，代表这个数据触碰了我建立的name1规则。

其他Yara规则中标准的meta字段，hfish使用填空的形式进行收集，详情界面可以参见：

<img src="http://img.threatbook.cn/hfish/image-20220525204450737.png" alt="image-20220525204450737" style="zoom: 33%;" />



最后附上一篇较为全面的Yara检测规则语法参考

https://b1ackie.cn/2021/09/13/yara%E8%A7%84%E5%88%99%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%EF%BC%88%E4%B8%80%EF%BC%89/


# DesignPatterns

使用 golang 泛型实现《大话设计模式》

- 测试代码使用

```bash
go test -v .
```

## Ch1 简单工厂模式

[简单工厂模式](./ch1/simplyfactory.go)

## Ch2 策略模式

[策略模式](./ch2/strategy.go)

### 策略模式+简单工厂模式

[收费系统策略模式+简单工程模式](./ch2/cash.go)

### 装饰器模式

[装饰器模式](./ch6/person.go)

------------


## 第一章
### UML 类图

![](Pasted%20image%2020230525093807.png)

- 继承关系：实线+空心三角形（a 指向 b，a 继承于 b，b 是父类）
- 实现接口：虚线+空心三角形（a 指向接口 b，a 实现了接口 b）
- 依赖关系：虚线+箭头（a 指向 b 表示 a 依赖 b）

- 组合（合成）关系：实心菱形+实线+箭头（菱形向a 箭头向 b b是 a 的一部分，“合成（Composition，也有翻译成‘组合’的）是一种强的‘拥有’关系，体现了严格的部分和整体的关系，部分和整体的生命周期一样）
	- 注意到合成关系的连线两端还有一个数字‘1’和数字‘2’，这被称为基数。表明这一端的类可以有几个实例
- 聚合关系）（Aggregation）关系：空心的菱形+实线+箭头（a指向 b，聚合表示一种弱的‘拥有’关系，体现的是A对象可以包含B对象，但B对象不是A对象的一部分,每只大雁都是属于一个雁群，一个雁群可以有多只大雁)

- 关联关系：实线+箭头 （a 指向 b，a 知道 b）


## 第三章 单一职责原则
- 单一职责原则：就一个类而言，应该仅有一个引起它变化的原因
- 一个类型承担的职责越多，它被复用的可能性就越小
- 软件设计真正要做的许多内容，就是发现职责并把那些职责相互分离,
- 如果你能够想到多于一个的动机去改变一个类，那么这个类就具有多于一个的职责[ASD]，就应该考虑类的职责分离。”



## 第四章 开放封闭原则
- 对扩展开放，对修改封闭-
- 但我们却可以在发生小变化时，就及早去想办法应对发生更大变化的可能。等到变化发生时立即采取行动[ASD]。正所谓，同一地方，摔第一跤不是你的错，再次在此摔跤就是你的不对了。
- 开放-封闭原则是面向对象设计的核心所在。遵循这个原则可以带来面向对象技术所声称的巨大好处，也就是可维护、可扩展、可复用、灵活性好。开发人员应该仅对程序中呈现出频繁变化的那些部分做出抽象，然而，对于应用程序中的每个部分都刻意地进行抽象同样不是一个好主意。拒绝不成熟的抽象和抽象本身一样重要[ASD]。切记，切记。


## 第五章 依赖倒置
依赖倒转原则
- A．高层模块不应该依赖低层模块。两个都应该依赖抽象。
	- 比如我们做的项目大多要访问数据库，所以我们就把访问数据库的代码写成了函数，每次做新项目时就去调用这些函数。这也就叫做高层模块依赖低层模块。
	- 就跟测试工具一样，不应该在被测试代码里面注入自定义的东西（因为需要更改）
- B．抽象不应该依赖细节。细节应该依赖抽象。[ASD]

里氏替换原则
个软件实体如果使用的是一个父类的话，那么一定适用于其子类，而且它察觉不出父类对象和子类对象的区别。也就是说，在软件里面，把父类都替换成它的子类，程序的行为没有变化，简单地说，子类型必须能够替换掉它们的父类型[ASD]。
![](Pasted%20image%2020230526160606.png)


## 装饰模式
装饰模式（Decorator），动态地给一个对象添加一些额外的职责，就增加功能来说，装饰模式比生成子类更为灵活。[DP]

![](Pasted%20image%2020230529111810.png)


Component是定义一个对象接口，可以给这些对象动态地添加职责。ConcreteComponent是定义了一个具体的对象，也可以给这个对象添加一些职责。
Decorator，装饰抽象类，继承了Component，从外类来扩展Component类的功能，但对于Component来说，是无需知道Decorator的存在的。
至于ConcreteDecorator就是具体的装饰对象，起到给Component添加职责的功能[DPE]。
Decorator is Part of component 


## 工厂模式
工厂方法模式实现时，客户端需要决定实例化哪一个工厂来实现运算类，选择判断的问题还是存在的，也就是说，工厂方法把简单工厂的内部逻辑判断移到了客户端代码来进行。你想要加功能，本来是改工厂类的，而现在是修改客户端！”


## 原型模式
原型模式（Prototype），用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。[DP]
原型模式其实就是从一个对象再创建另外一个可定制的对象，而且不需知道任何创建的细节


## 模板方法
模板方法模式，定义一个操作中的算法的骨架，而将一些步骤延迟到子类中。模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。

```go
//金庸小说考题试卷 class TestPaper { public void TestQuestion1() { Console.WriteLine(" 杨过得到，后来给了郭靖，炼成倚天剑、屠龙刀的玄铁可能是[ ] a.球磨铸铁 b. 马口铁 c.高速合金钢 d.碳素纤维 "); } public void TestQuestion2() { Console.WriteLine(" 杨过、程英、陆无双铲除了情花，造成[ ] a.使这种植物不再害人 b.使一种珍 稀物种灭绝 c.破坏了那个生物圈的生态平衡 d.造成该地区沙漠化"); } public void TestQuestion3() { Console.WriteLine(" 蓝凤凰致使华山师徒、桃谷六仙呕吐不止，如果你是大夫，会给他们开什么药[ ] a.阿司匹林 b.牛黄解毒片 c.氟哌酸 d.让他们喝大量的生牛奶 e.以上全不对"); } }
```

## 迪米特法则
迪米特法则（LoD），如果两个类不必彼此直接通信，那么这两个类就不应当发生直接的相互作用。如果其中一个类需要调用另一个类的某一个方法的话，可以通过第三者转发这个调用。[J&DP]


## 外观模式
外观模式（Facade），为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。[DP]
首先，在设计初期阶段，应该要有意识的将不同的两个层分离，比如经典的三层架构，就需要考虑在数据访问层和业务逻辑层、业务逻辑层和表示层的层与层之间建立外观Facade，
![](Pasted%20image%2020230601163040.png)

## 建造者模式
![](Pasted%20image%2020230601165247.png)

- 建造者模式（Builder），将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。[DP]

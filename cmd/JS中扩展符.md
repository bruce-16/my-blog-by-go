title: JS中扩展符（spread/rest）的各种表现
categories: javascript
label: JS, ES6
---------------------------------------------


<blockquote class="blockquote-center">
每一天都以许下希望开始，以收获经验结束。
</blockquote>

# 深入 '...' 各种情况下的表现形式
在js中，合并多个对象是一种很常见的操作，在ES5的时候，没有一种很方便的语法来进行对象的合并。在ES6中引入了一个对象函数Object.assgin(source, [target])，再之后又引入了**对象的spread语法**。

<!-- more -->

如下：
```javascript
const cat = {
  legs: 4,
  sound: 'meow',
};
const dog = {
  ...cat,
  sound: 'woof'
};

/* 最后结果
dog => {
  legs: 4,
  sound: 'woof',
}
*/
```

## spread与属性的可枚举配置
在ES5及之后的规范中，对象的每一个属性都存在几个来描述该属性的属性。这些值用来描述对象的属性是否可写、可枚举和可配置的状态。这里只说可枚举属性，可枚举属性是一个bool值，代表对象属性是否可以被访问。我们可以使用`Object.keys()`来访问**自己的和可枚举属性**,也可以使用`for...in`语句来**遍历所有可枚举的属性**
如下：
```javascript
const person = {
  name: 'zachrey',
  age: 21,
};
Object.keys(person); // ['name', 'age']
console.log({...person});
/*
{
  name: 'zachrey',
  age: 21,
}
*/
```
所以`name` `age`是`person`对象中的可枚举属性，目前来说spread可以克隆所有的可枚举属性。
现在我们来给`person`定义一个不可以枚举的属性，再使用spread看能不能克隆出来。
```javascript
Object.defineProperty(person, 'sex', {
  enumerable: false,
  value: 'male',
});
console.log(person['sex']);// male

const clonePerson = {
  ...person,
};
console.log(Object.keys(person));// ['name', 'age']
console.log(clonePerson);
/*
{
  name: 'zachrey',
  age: 21,
}
*/
```
通过以上可以看出来：
1. 不可枚举属性，可以被访问，使用`console.log`打印出来。
2. `...`**并不能克隆**不可枚举属性。
3. `...`的表现形式与`Object.keys`的表现形式相同。

## spread与自身属性
对于一个js对象，它的属性可以是自己的也可能是原型链上的，接下来简单的实现以下继承。
```javascript
const personB = Object.create(person, {
  profession: {
    value: 'development',
    enumerable: true,
  }
});
console.log(personB.hasOwnProperty('profession')); // => true
console.log(personB.hasOwnProperty('name')); // => false
console.log(personB.hasOwnProperty('age')); // => false
```
如上，只有`profession`是属于`personB`自身的。
**Object spread从自己的源属性中进行复制的时候，会忽略继承的属性。**
如下：
```javascript
const cloneB = { ...personB };
console.log(cloneB); // => { profession: 'development' }
```
> Object spread可以从源对象中复制自己的和可枚举的属性。和Object.keys()相同。

## Object spread规则：最后属性获胜
*** 后者扩展属性覆盖具有相同键的早期属性 ***
如下：
```javascript
const cat = {
  sound: 'meow',
  legs: 4,
};

const dog = {
  ...cat,
  ...{
    sound: 'woof' // <----- 覆盖 cat.sound
  }
};
console.log(dog); // => { sound: 'woof', legs: 4 }

const anotherDog = {
  ...cat,
  sound: 'woof' // <---- Overwrites cat.sound
};
console.log(anotherDog); // => { sound: 'woof', legs: 4 }
```
## 浅拷贝
spread对值是复合类型的属性，只会拷贝它对该值得应用。
如下：
```javascript

const laptop = {
  name: 'MacBook Pro',
  screen: { size: 17, isRetina: true }
};
const laptopClone = { ...laptop };
console.log(laptop === laptopClone); // => false
console.log(laptop.screen === laptopClone.screen); // => true
```
首先比较`laptop === laptopClone`，其值是false。主对象被正确克隆。

然而，`laptop.screen === laptopClone.screen`值是true。这意味着，`laptop.screen`和`laptopClone.screen`引用相同的嵌套对象，但没有复制。

## 原型丢失

这里先声明一个类：
```javascript
class Game {
  constructor(name) {
    this.name = name;
  }
  getMessage() {
    return `I like ${this.name}!`;
  }
}
const doom = new Game('Doom');
console.log(doom instanceof Game); // => true console.log(doom.name); // => "Doom"
console.log(doom.getMessage()); // => "I like Doom!"
```
接下来我们使用spread克隆调用构造函数创建的`doom`实例：
```javascript
const doomClone = { ...doom };
console.log(doomClone instanceof Game); // => false
console.log(doomClone.name); // => "Doom"
console.log(doomClone.getMessage()); // => TypeError: doomClone.getMessage is not a function
```
`...doom`只将自己的属性name复制到`doomClone`而已。doomClone是一个普通的JavaScript对象，其原型是`Object.prototype`，而不是`Game.prototype`，这是可以预期的。Object Spread**不保存源对象的原型**。
想要修复原型可以使用`Object.setPrototypeOf(doomClone, Game.prototype)`。
使用Object.assign()可以更合理的克隆doom：
```javascript
const doomFullClone = Object.assign(new Game(), doom);
console.log(doomFullClone instanceof Game); // => true
console.log(doomFullClone.name); // => "Doom"
console.log(doomFullClone.getMessage()); // => "I like Doom!"
```

## 传播undefined、null和基本类型
```javascript
const nothing = undefined;
const missingObject = null;
const two = 2;
console.log({ ...nothing }); // => { }
console.log({ ...missingObject }); // => { }
console.log({ ...two }); // => { }
```

## 总结

Object spread有一些规则要记住：

* 它从源对象中提取自己的和可枚举的属性
* 扩展的属性具有相同键的，后者会覆盖前者

与此同时，Object spread是简短而且富有表现力的，同时在嵌套对象上也能很好的工作，同时也保持更新的不变性。它可以轻松的实现对象克隆、合并和填充默认属性。

在结构性赋值中使用Object rest语法，可以收集剩余的属性。

实际上，Object rest和Object spread是JavaScript的重要补充。

>[原文链接](https://www.w3cplus.com/javascript/how-three-dots-changed-javascript-object-rest-spread-properties.html "原文链接")

!['干巴爹'](/uploads/ganbadie.jpg)
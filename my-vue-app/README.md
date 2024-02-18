# my-vue-app

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

安装 Vue Sweetalert2
首先，需要安装 vue-sweetalert2。在项目目录中，运行以下命令：

```
npm install --save vue-sweetalert2
```


## 使用Vuex实现状态共享

对于需要在多个组件之间共享状态的复杂应用，Vuex是Vue官方推荐的状态管理库。它提供了一种集中管理应用所有组件的状态的机制。

实现步骤：

在项目中安装和配置Vuex。
在Vuex的store中定义一个状态变量（比如isLoggedIn）和修改这个状态的mutation或action。
在login.vue登录成功后，提交一个mutation或dispatch一个action来更新isLoggedIn状态。
在App.vue中通过计算属性来监听这个状态的变化，并据此显示或隐藏登录/注册按钮。

示例代码
这里提供一个使用Vuex的简单示例：

1. 安装Vuex：
如果还没有安装，通过npm或yarn安装Vuex：

```css
npm install vuex --save
```
或

```csharp
yarn add vuex
```
1.设置Vuex Store：
在项目中创建一个store.js文件，并设置初始状态和mutations：

```javascript
// store.js
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    isLoggedIn: false,
  },
  mutations: {
    setLoggedIn(state, status) {
      state.isLoggedIn = status;
    },
  },
});
```

2.在main.js中引入并使用store：

```javascript
import Vue from 'vue';
import App from './App.vue';
import store from './store'; // 引入store

new Vue({
  store, // 使用store
  render: h => h(App),
}).$mount('#app');
```
3.在login.vue中提交mutation更新状态：

```javascript
// 在登录成功的处理函数中
this.$store.commit('setLoggedIn', true);
```

4.在App.vue中使用计算属性来动态显示内容：

```javascript
computed: {
  isLoggedIn() {
    return this.$store.state.isLoggedIn;
  }
}
```



## 使用 bootstrap-vue-3实现图标

### Usage
In your main.js or entry file:

```javascript
import { createApp } from 'vue'
import App from './App.vue'
import { BootstrapVue3 } from 'bootstrap-vue-3'

// Optionally import Bootstrap and BootstrapVue CSS
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue-3/dist/bootstrap-vue-3.css'

const app = createApp(App)

app.use(BootstrapVue3)

app.mount('#app')

```

### Tailwind CSS
Tailwind CSS, known for its utility-first approach, allows for highly customizable designs without stepping away from your markup.

#### Installation
```bash
npm install tailwindcss
```
Follow the official Tailwind CSS installation guide for Vue to set it up correctly in your project.

### Material Design (Using Vuetify with Vue 3)

Vuetify offers a vast collection of Material Design components for Vue.js.

Installation
For Vue 3 (Vuetify 3 is compatible with Vue 3):

```bash
npm install vuetify@next
```

#### Usage
In your main.js:

```javascript
import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'

const app = createApp(App)

app.use(vuetify)

app.mount('#app')
```
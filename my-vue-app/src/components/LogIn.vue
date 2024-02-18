<template>
  <div id="log-in">
    <form @submit.prevent="login">
      <input v-model="phone" type="text" placeholder="Phone" />
      <input v-model="password" type="password" placeholder="Password" />
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      phone: "",
      password: "",
    };
  },
  methods: {
    async login() {
      try {
        const response = await fetch(
          "http://localhost:8083/api/v1/login/common", // 替换为你的后端登录接口地址
          {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              phone: this.phone,
              password: this.password,
            }),
          }
        );
        const data = await response.json();
        if (response.ok) {
            this.$swal({
              title: "登录成功",
              icon: "success",
            });
          // 登录成功处理
          console.log("Login success:", data);
          // 更新Vuex中的登录状态
          this.$store.commit('setLoggedIn', true);

          this.$router.push('/dashboard');
          // 这里可以添加跳转到主页或其他页面的代码
        } else {
          this.$swal({
            title: "登录失败",
            text: data.message,
            icon: "error",
          });
          console.error("Login error:", data);
        }
      } catch (error) {
        this.$swal({
          title: "错误",
          text: "登录过程中发生错误。",
          icon: "error",
        });
        console.error("Login error:", error);
      }
    },
  },
};
</script>

<style scoped>
#log-in {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f2f2f2;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
    "Helvetica Neue", Arial, sans-serif;
  color: #333;
  text-align: center;
}

form {
  background: white;
  padding: 40px;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  width: 300px;
}

input {
  width: 100%;
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ddd;
  border-radius: 5px;
  box-sizing: border-box;
  transition: border-color 0.3s;
}

input:focus {
  border-color: #007aff;
  outline: none;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #007aff;
  border: none;
  border-radius: 5px;
  color: white;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #0056b3;
}

button:active {
  background-color: #003d7a;
}
</style>


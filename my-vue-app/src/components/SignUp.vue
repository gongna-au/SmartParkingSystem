<template>
    <div id="sign-up">
      <form @submit.prevent="signup">
        <input v-model="phone" type="text" placeholder="Phone" />
        <input v-model="password" type="text" placeholder="Password" />
        <button type="submit">Sign Up</button>
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
      async signup() {
        try {
          const response = await fetch(
            "http://localhost:8083/api/v1/sign/common",
            {
              // 替换为你的后端接口地址
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
              title: "注册成功",
              icon: "success",
            });
            
            this.$router.push('/login');
            console.log("Success:", data);
          } else {
            this.$swal({
              title: "注册失败",
              text: data.message,
              icon: "error",
            });
            console.error("Signup error:", data);
          }
        } catch (error) {
          alert("注册过程中发生错误。");
          console.error("Signup error:", error);
        }
      },
    },
  };
  </script>
  
  <style>
  #sign-up {
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
  
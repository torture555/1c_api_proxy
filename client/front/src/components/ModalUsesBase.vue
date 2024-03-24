<script setup>
import Input from "@/components/InputForForms.vue";
import {GetEnumColor} from "@/common.js";
</script>

<script>

import {URL_BACKEND} from "@/constants.js";
import {GetEnumColor, GetEnumValue} from "@/common.js";

export default {
  props: {
    baseName: {
      type: String
    }
  },
  data() {
    setInterval(this.updateStatus, 5000)
    this.getBase()
    return {
      labels: [],
      errorComment: "",
      status: false
    }
  },
  computed: {
    setStyle() {
      let color = GetEnumColor("StatusConnection", this.status)
      if (color === undefined) {
        color = "#D3D3D3"
      }
      return "background-color: " + color
    }
  },
  methods: {
    async updateStatus() {
      let resp = await fetch(URL_BACKEND + GetEnumValue("Methods", "StatusInfobase"), {
        headers: {
          Infobase: this.getValue('Имя базы'),
        }
      });
      let statusJson = await resp.json();
      if (statusJson !== undefined || statusJson !== Promise) {
        this.status = statusJson.result;
      }
    },
    async getBase() {

      let resp = await fetch(URL_BACKEND + GetEnumValue("Methods", "GetInfobases"))
      let bases = await resp.json()
      let base = bases.find((value) => value.name === this.baseName)
      if (base === undefined) {
        return
      }

      this.labels = [
        {
          name: 'Имя базы',
          value: base.name
        },
        {
          name: 'URL',
          value: base.URL
        },
        {
          name: 'Логин',
          value: base.login
        },
        {
          name: 'Пароль',
          value: ""
        },
      ]
    },
    setValue(name, value) {
      let pair = this.labels.find((pair) => (pair.name === name))
      if (pair !== undefined) {
        pair.value = value
      }
    },
    async updateParams() {
      await fetch(URL_BACKEND + GetEnumValue("Methods", "EditInfobase"), {
        method: "POST",
        body: JSON.stringify(
            {
              name: this.getValue('Имя базы'),
              URL: this.getValue('URL'),
              login: this.getValue('Логин'),
              password: this.getValue('Пароль'),
            })}).then((response) => {
            console.log(response.status)
            if (response.status >= 400) {
              this.error = true
              switch (response.status) {
                case 400:
                  this.errorComment = "Не правильно сформирован запрос"
                      break
                case 500:
                  this.errorComment = "Произошла ошибка на сервере"
                      break
              }
            } else {
              this.$emit("closeModal")
            }
          }
      )
    },
    getValue(name) {
      let pair = this.labels.find((pair) => (pair.name === name))
      if (pair !== undefined) {
        return pair.value
      } else {
        return ""
      }
    }
  }
}

</script>

<template>
<div class="modal-absolute">
  <div class="modal-relative">

    <div class="modal-window">
      <div class="modal-window-header">
        <h2>ИНФОРМАЦИОННАЯ БАЗА</h2>
        <button @click="$emit('closeModal')">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="white" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
      <div class="modal-status-base">
        <h3>Статус</h3>
        <div class="modal-status-base-div">
          <div class="modal-status-circle" :style="setStyle"></div>
          <h4>{{GetEnumValue("StatusConnection", status)}}</h4>
        </div>
      </div>
      <Input
        v-for="label in labels"
        :key="label.index"
        :name-of-input="label"
        @setValue="setValue"
      />
      <label v-if="error" style="color: red; text-align: center">
        {{errorComment}}
      </label>
      <div class="modal-footer">
        <button class="btn-refresh">
          <svg @click="updateStatus" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke=" rgba(121, 112, 233, 1)" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
          </svg>
        </button>
        <button
            @click="updateParams"
            class="btn-save">
          Сохранить
        </button>
      </div>
    </div>

  </div>
</div>
</template>

<style scoped>

.modal-absolute {

  width: 100%;
  height: 100%;

  position: absolute;

  z-index: 5;

}

.modal-relative {

  width: 100%;
  height: 100%;

  display: flex;
  justify-content: center;
  align-items: center;

  background-color: rgba(13, 11, 11, 0.9)

}

.modal-window {

  position: absolute;

  width: 21vw;

  background: #282828;

  z-index: 10;

  padding: 1.2%;

  border-radius: 8px;
}

.modal-window-header {
  width: 100%;
  height: 2rem;

  display: flex;
  justify-content: space-between;
  align-items: center;

  margin-bottom: 1rem;
}

.modal-window-header h2 {

  font-weight: 600;
  font-size: 1rem;
  color: white;
}

.modal-window-header button {

  height: 100%;
  aspect-ratio: 1 / 1;

}

.modal-status-base {

  width: 100%;
  height: 3.5rem;

  margin-bottom: 2rem;

}

.modal-status-base h3{

  width: 100%;
  height: 1rem;

  color: white;

  margin-bottom: 0.5rem;

}

.modal-status-base-div {
  width: 100%;
  height: 2rem;

  color: white;

  border-bottom: gray 1px solid;

  display: flex;
  justify-content: left;
  align-items: center;
}

.modal-status-circle {
  height: 27%;
  aspect-ratio: 1 / 1;
  border-radius: 100%;

  margin-right: 3%;

}

.modal-footer {

  width: 100%;
  height: 2.5rem;

  display: flex;
  justify-content: space-between;
  align-items: center;

}

.btn-refresh {
  height: 80%;
  aspect-ratio: 1 / 1;

}

.btn-save {
  height: 100%;
  aspect-ratio: 3.9 / 1;

  background: linear-gradient(70deg, rgba(121, 112, 233, 1), rgba(32, 161, 251, 1));

  border-radius: 7px;

  font-weight: 600;

}

</style>
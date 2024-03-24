<script setup>
import Input from "@/components/InputForForms.vue";
</script>

<script>

import {URL_BACKEND} from "@/constants.js";
import {GetEnumValue} from "@/common.js";

export default {
  data() {
    this.loadLabels()
    let labels = []
    return {
      labels: labels,
      error: false,
      errorComment: ""
    }
  },
  methods: {
    loadLabels() {
      this.fetchParams().then((value) => {
        this.labels = [
          {
            name: 'Хост',
            value: value.host
          },
          {
            name: 'Порт',
            value: value.port
          },
          {
            name: 'Логин',
            value: value.login
          },
          {
            name: 'Имя базы данных',
            value: value.DBName
          },
          {
            name: 'Пароль',
            value: ""
          },
        ]
      })

    },
    async fetchParams() {
      let params = await fetch(URL_BACKEND + GetEnumValue("Methods", "GetDBParams"))
      return await params.json()
    },
    setValue(name, value) {
      let pair = this.labels.find((pair) => (pair.name === name))
      if (pair !== undefined) {
        pair.value = value
      }
    },
    async updateParams() {
      await fetch(URL_BACKEND + GetEnumValue("Methods", "SetDBParams"), {
        method: "POST",
        body: JSON.stringify(
            {
              host: this.getValue('Хост'),
              port: this.getValue('Порт'),
              login: this.getValue('Логин'),
              password: this.getValue('Пароль'),
              DBName: this.getValue('Имя базы данных')
            })}).then((response) => {
        console.log(response.status)
            if (response.status >= 400) {
              this.error = true
              switch (response.status) {
                case 400:
                  this.errorComment = "Не правильно сформирован запрос"
                case 500:
                  this.errorComment = "Произошла ошибка на сервере"
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
  },
  emits: ["closeModal"]
}

</script>

<template>
  <div class="modal-absolute">
    <div class="modal-relative">

      <div class="modal-window">
        <div class="modal-window-header">
          <h2>ПРОВЕРКА БД</h2>
          <button @click="$emit('closeModal')">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="white" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" />
            </svg>
          </button>
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
          <button
              @click="updateParams"
              class="btn-save"
          >
            Изменить параметры
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

  background-color: lawngreen;

  margin-right: 3%;

}

.modal-footer {

  width: 100%;
  height: 2.5rem;

  display: flex;
  justify-content: center;
  align-items: center;

}

.btn-refresh {
  height: 80%;
  aspect-ratio: 1 / 1;

}

.btn-save {
  height: 100%;
  aspect-ratio: 7 / 1;

  background: linear-gradient(70deg, rgba(121, 112, 233, 1), rgba(32, 161, 251, 1));

  border-radius: 7px;

  font-weight: 600;

}

</style>
<script setup>
import LogItem from "@/components/Log/LogItem.vue";
</script>

<script>

import {GetEnumValue} from "@/common.js";
import {URL_BACKEND} from "@/constants.js";

export default {
  data() {
    this.updateStatus(true)
    setInterval(this.updateStatus, 3000)
    setInterval(this.getLogs, 10000)

    return {
      isConnected: false,
      logs: []
    }

  },
  methods: {
    async updateStatus(getLogs = false) {
      let res = await fetch(URL_BACKEND + GetEnumValue("Methods", "GetDBStatus"))
      res = await res.json()
      this.isConnected = res.status
      if (getLogs) {
        await this.getLogs()
      }
    },
    async getLogs() {
      if (this.isConnected) {
        let res = await fetch(URL_BACKEND + GetEnumValue("Methods", "GetLogs"))
        res = await res.json()
        console.log(res)
        this.logs = res.l
      }
    }
  },


}

</script>

<template>
  <div class="info-data-page-container">
    <div class="info-data-page-content">
      <div class="info-data-header">
        <h1>ЛОГ</h1>
      </div>
      <div class="statuses-div">
        <div class="btn-new-data">
          <button
              @click="$emit('openModalLog')"
          >Параметры БД</button>
        </div>
      </div>
      <div class="data-table">
        <div class="data-table-header">
          <h3 class="h3-status">Статус</h3>
          <h3 class="h3-name">Имя базы</h3>
          <h3 class="h3-date">Время</h3>
          <h3 class="h3-handler">Хендлер</h3>
          <h3 class="h3-comment">Комментарий</h3>
          <h3 class="h3-context">Контекст</h3>
        </div>
        <div class="data-table-border"></div>
        <div v-if="!isConnected" class="not-connected">
          <button style="cursor: pointer">
          <svg @click="updateStatus" style="height: 50px" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="green" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
          </svg>
          </button>
          <div>
            <label style="color: mintcream; text-align-all: center">
              База данных не подключена
            </label>
          </div>
        </div>
        <div v-if="isConnected" class="data-table-place-item">
          <LogItem
            v-for="raw in logs"
            :raw="raw"
          />
        </div>

      </div>

    </div>
  </div>
</template>

<style scoped>

h3 {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

.info-data-page-container {
  width: 52vw;
  height: 100vh;

  background-color: #282828;

  display: flex;
  justify-content: center;
  align-items: center;
}

.info-data-page-content {

  width: 90%;
  height: 80%;

}

.info-data-header {
  width: 100%;
  height: 9%;

  margin-bottom: 1%;

  font-size: 2.5rem;
  color: white;
  font-weight: 700;
}

.statuses-div {
  width: 100%;
  height: 6%;

  margin-bottom: 2%;

  display: flex;
  justify-content: left;
  align-items: center;
}

.btn-new-data {

  height: 100%;
  aspect-ratio: 5000/ 1137;

}

.not-connected {

  display: block;
  border: white solid 1px;
  border-radius: 10px;
  margin-top: 25%;
  width: 50%;
  text-align: center;
  margin-left: 25%;

}

.btn-new-data button {
  width: 100%;
  height: 100%;

  font-weight: 600;

  border-radius: 8px;

  background: linear-gradient(70deg, rgba(121, 112, 233, 1), rgba(32, 161, 251, 1));;
}

.data-table {

  width: 97%;
  height: 79%;

}

.data-table-border {

  width: 96%;

  border-top: 3px solid white;
}

.data-table-header {

  width: 97%;
  height: 5%;

  display: flex;
  justify-content: left;
  align-items: center;

  font-size: 1rem;
  font-weight: 600;
  color: white;

  margin-bottom: 1%;
}

.h3-status {
  height: 1rem;
  width: 12%;

  margin-right: 2%;
}

.h3-name {
  height: 1rem;
  width: 18%;

  margin-right: 2%;
}

.h3-date {
  height: 1rem;
  width: 15%;

  margin-right: 2%;
}

.h3-handler {
  height: 1rem;
  width: 12%;

  margin-right: 2%;
}

.h3-comment {
  height: 1rem;
  width: 17%;

  margin-right: 2%;
}

.h3-context {
  height: 1rem;
  width: 16%;
}

.data-table-place-item {

  height: 97%;
  width: 100%;

  overflow-y: auto;

}

.data-table-place-item::-webkit-scrollbar {
  width: 7px;
}

.data-table-place-item::-webkit-scrollbar-thumb {
  border-radius: 7px;
  background-color: gray;
}

</style>
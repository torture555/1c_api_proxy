<script>

import {URL_BACKEND} from "@/constants.js";
import {GetEnumColor, GetEnumValue} from "@/common.js";

export default {
  data() {
    this.updateStatus()
    setInterval(this.updateStatus, 10000)
    return {
      status: false
    }
  },
  props: {
    infobase: {
      type: Object
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
  methods:{
    GetEnumValue,
    async updateStatus() {
      let resp = await fetch(URL_BACKEND + GetEnumValue("Methods", "StatusInfobase"), {
        headers: {
          Infobase: this.infobase.name,
        }
      });
      if (resp.status === 408) {
        this.status = false
        return
      }
      let statusJson = await resp.json();
      if (statusJson !== undefined || statusJson !== Promise) {
        this.status = statusJson.result;
      }
    },
    async reloadPing() {
      await fetch(URL_BACKEND + GetEnumValue("Methods", "ReloadConnection"), {
        headers: {
          Infobase: this.infobase.name,
        }
      })
    }
  },
  emits: ["editBase"]
}

</script>

<template>
<div class="info-data-item">
      <div class="info-data-value">
          <div class="h3-status">
            <div class="status-circle" :style="setStyle"></div>
            <p>{{ GetEnumValue("StatusConnection", this.status) }}</p>
          </div>
          <p class="h3-name" :title="infobase.name">{{infobase.name}}</p>
          <p class="h3-url" :title="infobase.URL">{{ infobase.URL }}</p>
          <p class="h3-login" :title="infobase.login">{{ infobase.login }}</p>
      </div>
      <div class="info-data-item-btns">
        <button>
          <svg @click="reloadPing" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1" stroke="rgba(79, 135, 242, 1)" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
          </svg>
        </button>
        <button
        @click="$emit('editBase', infobase.name)">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1" stroke="rgba(79, 135, 242, 1)" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10" />
          </svg>
        </button>
      </div>

</div>
  <div class="info-data-border-bottom"></div>
</template>

<style scoped>

p {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.info-data-item {
  width: 100%;
  height: 7%;

  display: flex;
  justify-content: left;
  align-items: center;
}

.info-data-border-bottom {
  width: 97%;

  border-bottom: 1px solid gray;

  display: flex;
  justify-content: left;
  align-items: center;
}

.info-data-value {
  width: 86%;
  height: 100%;

  display: flex;
  justify-content: left;
  align-items: center;

  font-size: 0.8rem;

  color: white;
}

.h3-status {
  height: 1rem;
  width: 25%;

  margin-right: 2%;

  display: flex;
  justify-content: left;
  align-items: center;
}

.status-circle {
  height: 65%;
  aspect-ratio: 1 / 1;

  border-radius: 100%;

  margin: 0 7% 0 0;
}

.h3-name {
  height: 1rem;
  width: 24%;

  margin-right: 2%;
}

.h3-url {
  height: 1rem;
  width: 25%;
  margin-right: 2%;
}

.h3-login {
  height: 1rem;
  width: 20%;

}

.info-data-item-btns {

  width: 10%;
  height: 100%;

  display: flex;
  justify-content: space-between;
  align-items: center;
}

.info-data-item-btns button {
  height: 50%;
  aspect-ratio: 1 / 1;
}

.info-data-item-btns button:first-child {
  margin-left: 25%;
}



</style>
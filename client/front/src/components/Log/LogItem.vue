<script setup>

import {GetEnumValue} from "../../common.js";
</script>

<script>

import {GetEnumColor} from "@/common.js";

export default {
  data() {

    return {
      date: new Date(this.raw.date)
    }
  },
  props: {
    raw: {
      type: Object
    }
  },
  computed: {
    setStyle() {
      let color = GetEnumColor("LogsLevel", this.raw.level)
      if (color === undefined) {
        color = "#D3D3D3"
      }
      return "background-color: " + color
    }
  },
}

</script>

<template>
  <div class="info-data-item">
    <div class="info-data-value">
      <div class="h3-status">
        <div class="status-circle" :style="setStyle"></div>
        <p>{{ GetEnumValue("LogsLevel", raw.level) }}</p>
      </div>
      <p class="h3-name" :title="raw.base_name">{{raw.base_name}}</p>
      <p class="h3-date" :title="date.toLocaleTimeString()">{{date.toLocaleTimeString()}}</p>
      <p class="h3-handler" :title="raw.handler">{{raw.handler}}</p>
      <p class="h3-comment" :title="raw.comment">{{raw.comment}}</p>
      <p class="h3-context" :title="raw.context">{{raw.context}}</p>
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
  width: 97%;
  height: 100%;
  column-count: 6;
  display: flex;
  justify-content: left;
  align-items: center;

  font-size: 0.8rem;

  color: white;
}

.h3-status {
  height: 1rem;
  width: 12%;

  margin-right: 2%;

  display: flex;
  justify-content: left;
  align-items: center;

  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;

}

.status-circle {
  height: 65%;
  aspect-ratio: 1 / 1;

  border-radius: 100%;

  margin: 0 7% 0 0;
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



</style>
<script setup>
import LogItem from "@/components/Log/LogItem.vue";
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
          <h3 class="h3-date">Дата</h3>
          <h3 class="h3-handler">Хендлер</h3>
          <h3 class="h3-comment">Комментарий</h3>
          <h3 class="h3-context">Контекст</h3>
        </div>
        <div class="data-table-border"></div>
        <div class="data-table-place-item">
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>
          <LogItem/>


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
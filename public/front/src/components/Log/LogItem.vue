<script setup>

</script>

<template>
  <div class="info-data-item">
    <div class="info-data-value">
      <div class="h3-status">
        <div class="status-circle"></div>
        <p>Error</p>
      </div>
      <p class="h3-name">Имя базы какой-то не знаю ла-ла-ла</p>
      <p class="h3-date">Дата какой-то не знаю ла-ла-ла</p>
      <p class="h3-handler">Хендлер какой-то не знаю ла-ла-ла</p>
      <p class="h3-comment">Комментарий какой-то не знаю ла-ла-ла</p>
      <p class="h3-context">Контекст какой-то не знаю ла-ла-ла</p>
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

  background-color: yellow;
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
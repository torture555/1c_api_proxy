<script setup>
import Input from "@/components/InputForForms.vue";
</script>

<script>

export default {
  data() {
    return {
      labels: [
        {
          name: 'Имя базы'
        },
        {
          name: 'URL'
        },
        {
          name: 'Логин'
        },
        {
          name: 'Пароль'
        },
      ]
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
          <div class="modal-status-circle"></div>
          <h4>Есть соединение</h4>
        </div>
      </div>
      <Input
        v-for="label in labels"
        :key="label.index"
        :name-of-input="label"
      />
      <div class="modal-footer">
        <button class="btn-refresh">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke=" rgba(121, 112, 233, 1)" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
          </svg>
        </button>
        <button class="btn-save">Сохранить</button>

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
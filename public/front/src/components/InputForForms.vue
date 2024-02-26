<script>

export default {
  props: {
    nameOfInput: Object,
  }
}

</script>

<template>

  <div class="input">
    <h4>{{nameOfInput.name}}</h4>
    <input placeholder="Введите значение">
  </div>

</template>

<style scoped>

.input {

  width: 100%;
  height: 3.5rem;

  margin-bottom: 2rem;

}

.input h4 {

  width: 100%;
  height: 1rem;

  color: white;

  margin-bottom: 0.5rem;

}

.input input {

  width: 100%;
  height: 2rem;

  background-color: none;
  border-left: 0;
  border-top: 0;
  border-right: 0;

  border-bottom: gray 1px solid;

  color: white;

}

</style>
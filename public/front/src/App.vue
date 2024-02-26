<script setup>
import InfoDataPage from "@/components/DataBase/InfoDataPage.vue";
import LogPage from "@/components/Log/LogPage.vue";
import ModalUsesBase from "@/components/ModalUsesBase.vue";
import ModalNewBase from "@/components/ModalNewBase.vue";
import ModalLogParameters from "@/components/ModalLogParameters.vue";
</script>

<script>

export default {
  data () {
    return {
      isUseBaseModal: false,
      isNewBaseModal: false,
      isLogParametersModal: false,

    }
  }
}

</script>

<template>
  <div class="page-structure">
    <InfoDataPage
      @openEditForm="isUseBaseModal = true"
      @openNewBase="isNewBaseModal = true"/>
    <LogPage
      @openModalLog="isLogParametersModal = true"/>
    <ModalUsesBase
      v-if="isUseBaseModal"
      @closeModal="isUseBaseModal = false"/>
    <ModalNewBase
        v-if="isNewBaseModal"
        @closeModal="isNewBaseModal = false"/>
    <ModalLogParameters
        v-if="isLogParametersModal"
        @closeModal="isLogParametersModal = false"/>

  </div>



</template>

<style scoped>
.page-structure {
  display: flex;
  font-family: Montserrat;
}

::-webkit-scrollbar {
  width: 0;
}

</style>

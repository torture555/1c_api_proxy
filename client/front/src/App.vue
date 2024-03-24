<script setup>
import InfoDataPage from "@/components/DataBase/InfoDataPage.vue";
import LogPage from "@/components/Log/LogPage.vue";
import ModalUsesBase from "@/components/ModalUsesBase.vue";
import ModalNewBase from "@/components/ModalNewBase.vue";
import ModalLogParameters from "@/components/ModalLogParameters.vue";
</script>

<script>

export default {
  methods: {
    updatePage() {
      this.iotaInfoDataPage++
      this.iotaLogPage++
    },
    openEditForm(nameBase) {
      console.log(nameBase)
      this.EditNameBase = nameBase
      this.isUseBaseModal = true
    }
  },
  data () {
    return {
      isUseBaseModal: false,
      isNewBaseModal: false,
      isLogParametersModal: false,

      EditNameBase: "",

      iotaInfoDataPage: 0,
      iotaLogPage: 0,

    }
  }
}

</script>

<template>
  <div class="page-structure">
    <InfoDataPage :key="iotaInfoDataPage"
      @openEditForm="openEditForm"
      @openNewBase="isNewBaseModal = true"/>
    <LogPage :key="iotaLogPage"
      @openModalLog="isLogParametersModal = true"/>
    <ModalUsesBase
      v-if="isUseBaseModal"
      @closeModal="isUseBaseModal = false; updatePage()"
      :baseName="EditNameBase"
    />
    <ModalNewBase
        v-if="isNewBaseModal"
        @closeModal="isNewBaseModal = false; updatePage()"/>
    <ModalLogParameters
        v-if="isLogParametersModal"
        @closeModal="isLogParametersModal = false; updatePage()"/>

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

<template>
  <div class="workloads-search">
    <div class="search-div">
      <i class="input-icon fas fa-search"></i>
      <input
        type="text"
        name="search"
        class="text-input"
        placeholder="Search"
        autocomplete="off"
        v-model="searchTerm"
        icon-class="fa-home"
        @input="searchTermChanged"
      />
    </div>
    <div class="filters-div">
      <span class="namespace-label">Namespace:</span>
      <namespace-select></namespace-select>
    </div>
  </div>
</template>

<style scoped lang="scss">
@import "../../assets/scss/include";

.search-div {
  position: relative;
  // width: 97%;
  padding-left: 57px;
}

.workloads-search {
  position: relative;
  width: 100%;
  .text-input {
    height: 42px;
    border-radius: 7px;
    background-color: #f0f0fb;
    border: none;
    width: 100%;
    box-sizing: border-box;
    padding: 0 0 0 40px;
    font-family: "Ubuntu", sans-serif;
    font-weight: 300;
    font-size: 13px;
    color: #6a6c71;
    &::placeholder {
      color: #b6b9c3;
    }
    &:focus {
      outline: none;
      padding: 0 0 0 39px;
      color: #6a6c71;
      border: 1px solid #b6b9c3;
    }
  }
  .input-icon {
    position: absolute;
    left: 72px;
    top: 50%;
    color: #9095a4;
    transform: translateY(-50%);
    font-size: 14px;
  }
}

.filters-div {
  position: relative;
  width: 100%;
  padding: 10px 0;
}

.namespace-label {
  color: #2e395a;
  font-weight: 400;
  font-size: 15px;
  font-family: "Ubuntu", sans-serif;
  float: left;
  padding-right: 10px;
  transform: translateY(40%);
  -ms-transform: translateY(40%);
  -moz-transform: translateY(40%);
  -webkit-transform: translateY(40%);
  -o-transform: translateY(40%);
}
</style>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import NamespaceSelect from "./NamespaceSelect.vue";
import { Action } from "vuex-class";

@Component({
  components: {
    NamespaceSelect
  }
})
export default class WorkloadSearch extends Vue {
  public searchTerm: string = "";

  @Action("changeSearchTerm", { namespace: StoreNamespaces.workloads })
  protected changeSearchTerm: any;

  public searchTermChanged() {
    this.changeSearchTerm(this.searchTerm);
  }
}
</script>
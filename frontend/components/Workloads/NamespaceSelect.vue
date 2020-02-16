<template>
  <div class="namespace-select" v-show="namespaces.length">
    <div class="loading-overlay" v-if="loading">
      <div class="lds-ripple">
        <div></div>
        <div></div>
      </div>
    </div>
    <select
      v-model="namespace"
      class="namespace-input"
      @keyup.enter="selectNamespace"
      v-on:change="selectNamespace"
    >
      <option
        v-for="namespace in namespaces"
        v-bind:value="namespace"
        :key="namespace"
      >{{namespace}}</option>
    </select>
    <span class="namespace-loading">{{loading ? 'Loading...' : ''}}</span>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import { Action, Getter } from "vuex-class";

@Component({})
export default class NamespaceSelect extends Vue {
  public namespace: string = "";

  public loading: boolean = false;

  @Action("setCurrentNamespace", { namespace: StoreNamespaces.namespaces })
  public setCurrentNamespace: any;

  @Action("fetchNamespaces", { namespace: StoreNamespaces.namespaces })
  public fetchNamespaces: any;

  @Action("fetchWorkloads", { namespace: StoreNamespaces.workloads })
  public fetchWorkloads: any;

  @Getter("currentNamespace", { namespace: StoreNamespaces.namespaces })
  public currentNamespace: any;

  @Getter("namespaces", { namespace: StoreNamespaces.namespaces })
  public namespaces: any;

  public async mounted() {
    try {
      await this.fetchNamespaces();
    } catch (e) {
      throw "Error fetching namespaces";
    }
    this.namespace = this.currentNamespace;
    if (this.namespace) {
      this.selectNamespace();
    }
  }

  public async selectNamespace() {
    if (this.loading) {
      return;
    }
    this.setCurrentNamespace(this.namespace);

    this.loading = true;

    try {
      await this.fetchWorkloads(this.currentNamespace);
    } catch (e) {
      alert(
        "Error when retrieving workloads for namespace: " +
          this.currentNamespace
      );
    }

    this.loading = false;
  }
}
</script>

<style scoped lang="scss">
@import "../../assets/scss/include";

.loading-overlay {
  position: fixed; /* Sit on top of the page content */
  width: 100%; /* Full width (cover the whole page) */
  height: 100%; /* Full height (cover the whole page) */
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(
    255,
    255,
    255,
    0.4
  ); /* White background with opacity */
  z-index: 2; /* Specify a stack order in case you're using a different order for other elements */
  cursor: wait; /* Add a pointer on hover */
}
.lds-ripple {
  display: inline-block;
  position: relative;
  width: 100%;
  height: 100%;
  position: absolute;
  top: 50%;
  left: 50%;
}
.lds-ripple div {
  position: absolute;
  border: 4px solid rgb(50, 127, 243);
  opacity: 1;
  border-radius: 50%;
  animation: lds-ripple 1s cubic-bezier(0, 0.2, 0.8, 1) infinite;
}
.lds-ripple div:nth-child(2) {
  animation-delay: -0.5s;
}
@keyframes lds-ripple {
  0% {
    top: 36px;
    left: 36px;
    width: 0;
    height: 0;
    opacity: 1;
  }
  100% {
    top: 0px;
    left: 0px;
    width: 72px;
    height: 72px;
    opacity: 0;
  }
}
.namespace-select {
  .namespace-input {
    height: 30px;
    border-radius: 7px;
    background-color: #f0f0fb;
    border: none;
    width: 200px;
    box-sizing: border-box;
    padding: 0 0 0 15px;
    font-family: "Ubuntu", sans-serif;
    font-weight: 300;
    font-size: 13px;
    color: #6a6c71;
    &::placeholder {
      color: #b6b9c3;
    }
    &:hover {
      background: #f7f7fd;
    }
    &:focus {
      background: #f0f0fb;
      outline: none;
      padding: 0 0 0 14px;
      color: #6a6c71;
      border: 1px solid #b6b9c3;
    }
  }

  .namespace-loading {
    font-size: 12px;
    font-family: "Ubuntu", sans-serif;
    padding-left: 10px;
  }
}
</style>

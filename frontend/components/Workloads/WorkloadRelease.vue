<template>
  <span>
    <button
      class="release_button"
      :disabled="!getWorkload(workload).selected_tag.tag || getWorkload(workload).status == statuses.releasing"
      @click="release(workload)"
    >{{getWorkload(workload).status == statuses.releasing ? 'Relasing' : 'Release'}}</button>
  </span>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import { Getter, Action } from "vuex-class";
import { WorkloadStatuses } from "../../store/types/Workloads/WorkloadStatuses";

@Component({})
export default class WorkloadRelease extends Vue {
  @Prop() protected workload: any;

  protected statuses = WorkloadStatuses;

  @Action("releaseVersion", { namespace: StoreNamespaces.workloads })
  public releaseVersion: any;

  @Getter("getWorkload", { namespace: StoreNamespaces.workloads })
  public getWorkload: any;

  public release() {
    const workload = this.getWorkload(this.workload);
    const releaseData: any = {
      Workload: workload.id,
      Container: workload.container,
      Current: workload.image + ":" + workload.current_tag.tag,
      Target: workload.image + ":" + workload.selected_tag.tag
    };

    this.releaseVersion({ workload: this.workload, releaseData });
  }
}
</script>

<style lang="scss">
@import "../../assets/scss/include";
.release_button {
  background: #343d5d;
  padding: 8px;
  color: #fff;
  border-radius: 5px;
  font-family: "Ubuntu", sans-serif;
  font-size: 11px;
  border: none;
  transition-duration: 0.2s;
  &:hover {
    background: #5c6484;
    box-shadow: 0 3px 6px 0 rgba(0,0,0,0.4), 0 3px 8px 0 rgba(0,0,0,0.19);
    cursor: pointer;
  }
  &:active {
    background: #4a506b;
    box-shadow: 0 1px 2px 0 rgba(0,0,0,0.2), 0 1px 3px 0 rgba(0,0,0,0.19);
    cursor: pointer;
  }
  &:focus {
    outline: 0;
  }
  &:disabled {
    cursor: not-allowed;
    background: #f0f0fa;
    color: #b9b3b3;
    box-shadow: none;
  }
}
</style>
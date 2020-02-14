<template>
  <div class="button-div">
    <span>
      <button
        class="release-button"
        :disabled="!getWorkload(workload).selected_tag.tag || getWorkload(workload).status == statuses.releasing"
        @click="release(workload)"
      >{{getWorkload(workload).status == statuses.releasing ? 'Relasing' : 'Release'}}</button>
    </span>
    <span>
      <button
        class="more-button"
        :disabled="!getWorkload(workload).selected_tag.tag || getWorkload(workload).status == statuses.releasing"
        @click="toggleDropdown($event)"
      >
        <div class="icon-wrapper">
          &nbsp;&nbsp;&nbsp;
          <i class="fas fa-caret-down" aria-hidden="true"></i>
          <i class="fas fa-caret-up" aria-hidden="true"></i>
        </div>
      </button>
    </span>
  </div>
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

  public toggleDropdown(event) {
    console.log(event.currentTarget.classList);
    event.currentTarget.classList.toggle("is-active");
  }
}
</script>

<style lang="scss">
@import "../../assets/scss/include";

.button-div {
  font-size: 0;
  text-align: center;
}

.release-button {
  background: #343d5d;
  padding: 8px;
  color: #fff;
  border-radius: 5px 0 0 5px;
  font-family: "Ubuntu", sans-serif;
  font-size: 11px;
  transition-duration: 0.2s;
  border: solid;
  border-color: #5c6484;
  border-width: 0 0.5px 0 0;
  &:hover {
    background: #5c6484;
    box-shadow: 0 3px 6px 0 rgba(0, 0, 0, 0.4), 0 3px 8px 0 rgba(0, 0, 0, 0.19);
    cursor: pointer;
  }
  &:active {
    background: #4a506b;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.2), 0 1px 3px 0 rgba(0, 0, 0, 0.19);
    cursor: pointer;
  }
  &:focus {
    outline: 0;
  }
  &:disabled {
    cursor: not-allowed;
    background: #f0f0fa;
    color: #b9b3b3;
    border: solid;
    border-color: #343d5d10 #343d5d20 #343d5d10 #343d5d10;
    border-width: 0.5px;
    box-shadow: none;
  }
}

.more-button {
  background: #343d5d;
  padding: 8px 5px;
  color: #fff;
  border-radius: 0 5px 5px 0;
  font-family: "Ubuntu", sans-serif;
  font-size: 11px;
  border: none;
  transition-duration: 0.2s;
  .icon-wrapper {
    position: relative;
    height: 100%;
  }
  i.fas {
    position: absolute;
    left: 0;
    transition: opacity 0.3s;
    &.fa-caret-up {
      opacity: 0;
    }
  }
  &.is-active {
    i.fas {
      &.fa-caret-up {
        opacity: 1;
      }
      &.fa-caret-down {
        opacity: 0;
      }
    }
  }
  &:hover {
    background: #5c6484;
    box-shadow: 0 3px 6px 0 rgba(0, 0, 0, 0.4), 0 3px 8px 0 rgba(0, 0, 0, 0.19);
    cursor: pointer;
  }
  &:active {
    background: #4a506b;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.2), 0 1px 3px 0 rgba(0, 0, 0, 0.19);
    cursor: pointer;
  }
  &:focus {
    outline: 0;
  }
  &:disabled {
    cursor: not-allowed;
    background: #f0f0fa;
    color: #b9b3b3;
    border: solid;
    border-color: #343d5d10;
    border-width: 0.5px 0.5px 0.5px 0;
    box-shadow: none;
  }
}
</style>
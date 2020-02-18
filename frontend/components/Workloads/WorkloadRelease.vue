<template>
  <div class="button-div">
    <div class="button-release">
      <button
        class="button release"
        :disabled="!getWorkload(workload).selected_tag.tag || getWorkload(workload).status == statuses.releasing"
        @click="release(workload)"
      >{{getWorkload(workload).status == statuses.releasing ? 'Releasing' : 'Release'}}</button>
    </div>
    <div class="button-more" v-click-outside="clickOutsideDropdown">
      <!-- TODO: Add reset button to the "More" dialog to deselect the tag -->
      <button
        ref="more_btn"
        class="button more"
        :disabled="!getWorkload(workload).selected_tag.tag || getWorkload(workload).status == statuses.releasing"
        @click="toggleDropdown($event)"
      >
        <i class="fas fa-caret-down"></i>
      </button>
      <div class="button-more__content">
        <ul>
          <li>
            <i class="fas fa-star"></i> Promote
          </li>
          <!-- <li>Lorem ipsum dolor sit</li> -->
          <li>
            <i class="fas fa-trash-alt"></i> Clear Tag
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import { Getter, Action } from "vuex-class";
import { WorkloadStatuses } from "../../store/types/Workloads/WorkloadStatuses";

@Component({
  methods: {
    clickOutsideDropdown(){
      let moreBtn = this.$refs.more_btn as Element;
      moreBtn.classList.remove("is-active");
    }
  }
})
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

  public toggleDropdown(event: any) {
    event.currentTarget.classList.toggle("is-active");
  }
}
</script>

<style lang="scss">
@import "../../assets/scss/include";

.button-div {
  // TODO: make release button smaller when viewport too narrow.
  text-align: center;
  float: left;
  position: relative;
  .button-release {
    float: left;
    white-space: nowrap;
  }
  .button-more {
    white-space: nowrap;
    &__content {
      float: left;
      z-index: 999;
      position: absolute;
      right: 0;
      opacity: 0;
      overflow: hidden;
      transition: opacity 0.3s;
      visibility: hidden;

      font-family: "Ubuntu", sans-serif;
      font-size: 12px;
      background: #343d5de0;
      border-radius: 5px 0 5px 5px;
      box-shadow: 0 2px 6px 0 rgba(0, 0, 0, 0.3), 0 3px 9px 0 rgba(0, 0, 0, 0.1);
      ul {
        list-style-type: none;
        padding-left: 0;
        margin: 0;
        li {
          color: #ffffffd7;
          padding: 7px 15px 7px 10px;
          transition: 0.1s;
          cursor: pointer;
          user-select: none;
          i {
            padding-right: 5px;
          }
          &:hover {
            color: #fff;
            background: #576185e0;
            transition: 0.2s;
            i.fas {
              &.fa-star {
                color: #ffd000e0;
              }
              &.fa-trash-alt {
                color: #ff8a8ae0;
              }
            }
          }
        }
      }
    }
  }

  .button {
    background: #343d5d;
    color: #fff;
    font-family: "Ubuntu", sans-serif;
    font-size: 11px;
    transition-duration: 0.2s;
    &:hover {
      background: #5c6484;
      box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.4),
        0 1px 3px 0 rgba(0, 0, 0, 0.19);
      cursor: pointer;
    }
    &:active {
      background: #5c6484;
      box-shadow: inset 0 2px 5px 0 rgba(0, 0, 0, 0.4),
        inset 0 1px 3px 0 rgba(0, 0, 0, 0.25);
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
      box-shadow: none;
    }
  }

  .release {
    padding: 8px;
    border-radius: 5px 0 0 5px;
    border: solid;
    border-color: #5c6484;
    border-width: 0 0.5px 0 0;
    &:disabled {
      border-color: #343d5d10 #343d5d20 #343d5d10 #343d5d10;
      border-width: 0.5px;
    }
  }

  .more {
    padding: 8px 5px;
    border-radius: 0 5px 5px 0;
    border: none;
    i.fas {
      transition: 0.3s;
    }
    &.is-active {
      background: #5c6484;
      box-shadow: inset 0 1px 4px 0 rgba(0, 0, 0, 0.4),
        inset 0 1px 2px 0 rgba(0, 0, 0, 0.25);
      &:hover {
        box-shadow: inset 0 1px 4px 0 rgba(0, 0, 0, 0.4),
          inset 0 1px 2px 0 rgba(0, 0, 0, 0.25);
      }
      &:active {
        background: #5c6484;
        box-shadow: inset 0 2px 5px 0 rgba(0, 0, 0, 0.4),
          inset 0 1px 3px 0 rgba(0, 0, 0, 0.25);
        cursor: pointer;
      }
      i.fas {
        &.fa-caret-down {
          transform: rotate(180deg);
        }
      }

      + .button-more__content {
        opacity: 1;
        visibility: visible;
      }
    }
    &:disabled {
      border-color: #343d5d10;
      border-width: 0.5px 0.5px 0.5px 0;
    }
  }
}
</style>
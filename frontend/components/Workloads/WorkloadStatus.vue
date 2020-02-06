<template>
  <div>
    <div class="icon tooltip" :style="iconStyle">
      <span class="tooltiptext">{{ iconTooltip }}</span>
    </div>
  </div>
</template>

<style scoped lang="scss">
.icon {
  background-position: center center;
  background-size: contain;
  width: 35px;
  height: 35px;
  background-repeat: no-repeat;
}

/* Tooltip container */
.tooltip {
  position: relative;
  display: inline-block;
}

/* Tooltip text */
.tooltip .tooltiptext {
  visibility: hidden;
  width: 100px;
  background-color: #2f395b;
  color: #fff;
  text-align: center;
  padding: 5px 0;
  border-radius: 6px;

  /* Font Styling */
  font-family: "Ubuntu", sans-serif;
  font-weight: 300;
  font-size: 13px;

  /* Position the tooltip text */
  position: absolute;
  z-index: 1;
  bottom: 125%;
  left: 50%;
  margin-left: -60px;

  /* Fade in tooltip */
  opacity: 0;
  transition: opacity 0.3s;
}

/* Tooltip arrow */
.tooltip .tooltiptext::after {
  content: "";
  position: absolute;
  top: 100%;
  left: 50%;
  margin-left: -5px;
  border-width: 5px;
  border-style: solid;
  border-color: #555 transparent transparent transparent;
}

/* Show the tooltip text when you mouse over the tooltip container */
.tooltip:hover .tooltiptext {
  visibility: visible;
  opacity: 1;
}
</style>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { Getter } from "vuex-class";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import { WorkloadStatuses } from "../../store/types/Workloads/WorkloadStatuses";

@Component({})
export default class WorkloadStatus extends Vue {
  @Prop() protected workload: any;

  @Getter("getWorkload", { namespace: StoreNamespaces.workloads })
  public getWorkload: any;

  get iconStyle() {
    const status = this.getWorkload(this.workload).status;
    const imageName = status;
    const imageExtension = status == WorkloadStatuses.releasing ? "gif" : "png";
    return {
      backgroundImage: `url(${require(`@/assets/images/${imageName}.${imageExtension}`)})`
    };
  }

  get iconTooltip() {
    const status = this.getWorkload(this.workload).status;
    var statusLiteral = "";

    switch (status) {
      case "up_to_date":
        statusLiteral = "Up to Date";
        break;

      case "releasing":
        statusLiteral = "Releasing";
        break;

      case "behind":
        statusLiteral = "Behind";
        break;

      case "unknown":
        statusLiteral = "Unknown";
        break;

      case "release_failed":
        statusLiteral = "Release Failed";
        break;

      default:
        break;
    }

    return statusLiteral;
  }
}
</script>
<template>
  <div class="workloads-list">
    <span class="namespace-label">Namespace:</span>
    <namespace-select></namespace-select>
    <vue-good-table
      :columns="columns"
      :rows="workloads"
      styleClass="vgt-table striped"
      :search-options="{
        enabled: true,
        externalQuery: searchTerm
      }"
    >
      <template slot="table-row" slot-scope="props">
        <workload-release :workload="props.row" v-if="props.column.field == 'action'"></workload-release>
        <workload-status :workload="props.row" v-else-if="props.column.field == 'status'" />
        <workload-available-tags
          :options-prop="props.row.available_tags"
          :current-tag="props.row.current_tag"
          :workload="props.row"
          @input="tagChanged"
          v-else-if="props.column.field == 'available_tags'"
        />
        <div v-else>
          <div class="text-td">
            <span
              class="tooltiptext"
            >{{props.formattedRow[props.column.field]}}</span>
          </div>
        </div>
      </template>
    </vue-good-table>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import WorkloadAvailableTags from "./WorkloadAvailableTags.vue";
import WorkloadRelease from "./WorkloadRelease.vue";
import NamespaceSelect from "./NamespaceSelect.vue";
import WorkloadStatus from "./WorkloadStatus.vue";
import { namespace } from "vuex-class";
import { Getter, Action } from "vuex-class";
import { Workload } from "../../store/types/Workloads/Workload";
import { Tag } from "../../store/types/Workloads/Tag";

@Component({
  components: {
    WorkloadAvailableTags,
    NamespaceSelect,
    WorkloadRelease,
    WorkloadStatus
  }
})
export default class WorkloadsList extends Vue {
  public columns = [
    {
      label: "Workload",
      field: "workload",
      width: "20%"
    },
    {
      label: "Container",
      field: "container",
      width: "13%"
    },
    {
      label: "Image",
      field: "image",
      width: "32%"
    },
    {
      label: "Current tag",
      field: "current_tag.tag",
      width: "10%"
    },
    {
      label: "Status",
      field: "status",
      width: "5%"
    },
    {
      label: "Available tags",
      field: "available_tags",
      width: "15%",
      sortable: false
    }
  ];

  protected $env!: any;

  @Getter("message")
  protected message!: any;

  @Getter("workloads", { namespace: StoreNamespaces.workloads })
  protected workloads!: any;

  @Getter("searchTerm", { namespace: StoreNamespaces.workloads })
  protected searchTerm!: any;

  @Action("fetchNamespaces", { namespace: StoreNamespaces.namespaces })
  public fetchNamespaces: any;

  @Action("updateSelectedTag", { namespace: StoreNamespaces.workloads })
  public updateSelectedTag: any;

  public created() {
    if (this.$env.READ_ONLY != "true") {
      this.columns.push({
        label: "Action",
        field: "action",
        width: "5%"
      });
    }
  }

  public tagChanged(workload: Workload, tag: Tag) {
    this.updateSelectedTag({ workload, tag });
  }
}
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>

<style lang="scss">
@import "../../assets/scss/include";

/* Tooltip container */
.tooltip {
  position: relative;
  display: inline-block;
}

/* Tooltip text */
.tooltip .tooltiptext {
  visibility: hidden;
  width: 100%;
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
  margin-left: -50px;

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

/* Scrollbar Style */
/* width */
::-webkit-scrollbar {
  width: 10px;
}
/* Track */
::-webkit-scrollbar-track {
  display: none;
  background: #f1f1f1;
}
/* Handle */
::-webkit-scrollbar-thumb {
  background: #99a9c2;
  &:hover {
    background: #4b5677;
  }
  &:active {
    background: #303b5d;
  }
}

.text-td {
  overflow: hidden;
  width: -moz-available; /* WebKit-based browsers will ignore this. */
  width: -webkit-fill-available; /* Mozilla-based browsers will ignore this. */
  width: fill-available;
  text-overflow: ellipsis;
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

/* Hack to replace empty table text */
.vgt-text-disabled {
  visibility: hidden;
}
.vgt-text-disabled:after {
  content: "No workloads found in namespace.";
  visibility: visible;
  font-size: 16px;
  display: block;
  padding-bottom: 15px;
}

.workloads-list {
  height: calc(100% - 110px);
  min-height: 250px;
  padding: 15px 0;
  box-sizing: border-box;
  overflow-y: scroll;

  .vgt-responsive {
    overflow-x: initial;
  }
  .vgt-table {
    border: none;
    font-family: sans-serif;
    border-collapse: collapse;
    border-radius: 7px;
    width: 100%;
    &.striped {
      tr {
        &:nth-child(odd) {
          background: #f4f4fd;
        }
        &:nth-child(even) {
          background: #fafaff;
        }
        &:hover {
          background: #e6e6f7;
        }
      }
      td,
      th {
        border: none;
      }
      td {
        color: #3c5171;
        font-size: 14px;
        white-space: nowrap;
        font-family: "Ubuntu Mono", monospace;
        vertical-align: middle;
        padding-top: 2px;
        padding-bottom: 2px;
        max-width: 0;
      }
      th {
        color: #2e395a;
        font-weight: 400;
        font-size: 15px;
        background: #f0f0fa;
        font-family: "Ubuntu", sans-serif;
        &:hover {
          background: #e1e1f5;
        }
        &:active {
          background: #d4d4ee;
        }
      }
    }
  }
}
</style>
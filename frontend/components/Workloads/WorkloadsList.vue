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
        <span v-else>{{props.formattedRow[props.column.field]}}</span>
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
      field: "workload"
    },
    {
      label: "Container",
      field: "container"
    },
    {
      label: "Image",
      field: "image"
    },
    {
      label: "Current tag",
      field: "current_tag.tag"
    },
    {
      label: "Status",
      field: "status"
    },
    {
      label: "Available tags",
      field: "available_tags",
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
        field: "action"
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
<template>
  <div class="available-tag-select-wrapper">
    <multiselect
      v-model="value"
      :options="options"
      :placeholder="currentTag ? currentTag.tag : 'Select tag'"
      label="tag"
      track-by="tag"
      :allow-empty="false"
      deselect-label="Can't remove this value"
      selectedLabel=""
      @input="valueChanged"
    >
      <template slot="singleLabel" slot-scope="{ option }">
        <strong class="selected__tag">{{ option.tag }}</strong>
      </template>
      <template slot="option" slot-scope="props">
        <div class="option__desc">
          <span class="option__tag">{{ props.option.tag }}</span>
          <span class="option__date">
            <strong>{{ moment(props.option.date).fromNow() }}</strong>
          </span>
        </div>
      </template>
    </multiselect>
  </div>
</template>

<style lang="scss">
@import "../../assets/scss/include";

.available-tag-select-wrapper {
  width: 100%;
  .multiselect__tags {
    width: 100%;
    .multiselect__placeholder {
      width: 100%;
      overflow: hidden;
      text-overflow: ellipsis;
      font-family: "Ubuntu Mono", monospace;
      
    }
    .multiselect__single {
      width: 100%;
      overflow: hidden;
      text-overflow: ellipsis;
      .selected__tag {
        font-family: "Ubuntu Mono", monospace;
        font-size: 14px;
      }
    }
  }
  .multiselect__content-wrapper {
    min-width: max-content;
    right: 0;
  }
  .multiselect__option--highlight {
    background: #4b5677;
    &:after {
      background: #4b5677;
    }
  }
  .option__desc {
    .option__tag {
      font-family: "Ubuntu Mono", monospace;
      font-size: 14px;
    }
    .option__date {
      font-family: "Ubuntu", sans-serif;
      float: right;
      font-size: 14px;
    }
  }
}
</style>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import Multiselect from "vue-multiselect";
import moment from "moment";

@Component({
  components: {
    Multiselect
  }
})
export default class WorkloadAvailableTags extends Vue {
  @Prop() protected optionsProp: any;
  @Prop() protected currentTag: any;
  @Prop() protected workload: any;

  protected moment = moment;

  protected value: any = null;

  public updated() {
    this.refreshSelectedValue();
  }

  public refreshSelectedValue() {
    if (!this.value) {
      return;
    }
    const selectedTagFoundInProps = this.optionsProp.find(
      (o: any) => o.tag == this.value.tag
    );
    if (!selectedTagFoundInProps) {
      this.value = null;
    }
  }

  get options() {
    return this.optionsProp.map((option: any) => {
      option.$isDisabled = option.current;
      return option;
    });
  }

  public valueChanged() {
    this.$emit("input", this.workload, this.value);
  }
}
</script>
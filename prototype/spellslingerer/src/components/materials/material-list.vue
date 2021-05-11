<template>
  <div>
    <v-container v-if="oneline">
      <span class="pa-4 headline" v-if="label">{{ label }}</span>
      <v-row dense>
        <v-col>
          <v-slide-group show-arrows>
            <v-slide-item
              v-for="material in materials"
              :key="`${material.id}-${material.quantity}`"
              cols="3"
            >
              <material-list-item
                :readonly="readonly"
                v-bind="material"
                @updateQuantity="$emit('updateQuantity', $event)"
              />
            </v-slide-item>
          </v-slide-group>
        </v-col>
      </v-row>
    </v-container>
    <v-container v-else>
      <span class="pa-4 headline" v-if="label">{{ label }}</span>
      <v-row dense>
        <v-col
          v-for="material in materials"
          :key="`${material.id}-${material.quantity}`"
          cols="auto"
        >
          <material-list-item
            :readonly="readonly"
            v-bind="material"
            @updateQuantity="$emit('updateQuantity', $event)"
          />
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import MaterialListItem from './material-list-item.vue';

export default {
  name: 'material-list',
  components: {
    MaterialListItem,
  },
  props: {
    materials: {
      type: Array,
      required: true,
    },
    oneline: {
      type: Boolean,
      default: false,
    },
    readonly: {
      type: Boolean,
      default: false,
    },
    label: {
      type: String,
      default: '',
    },
  },
};
</script>

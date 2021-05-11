<template>
  <v-dialog v-model="open" max-width="600px">
    <v-card>
      <v-card-title class="headline">{{ title }}</v-card-title>
      <v-card-text>{{ body }}</v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn @click="clickCancel()">Cancel</v-btn>
        <v-btn :color="destructiveAction ? 'error' : 'success'" @click="clickConfirm()">OK</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: 'confirm-dialog',
  data() {
    return {
      awaitedResult: null,
    };
  },
  computed: {
    open: {
      get() {
        return this.value;
      },
      set(value) {
        this.$emit('input', value);
      },
    },
  },
  methods: {
    awaitResponse() {
      this.open = true;
      return new Promise((resolve) => {
        this.awaitedResult = resolve;
      });
    },
    clickConfirm() {
      if (this.awaitedResult) this.awaitedResult(true);
      this.$emit('confirmed');
      this.open = false;
    },
    clickCancel() {
      if (this.awaitedResult) this.awaitedResult(false);
      this.open = false;
    },
  },
  props: {
    value: {
      type: Boolean,
      required: true,
    },
    title: {
      type: String,
      default: '',
    },
    body: {
      type: String,
      default: '',
    },
    destructiveAction: {
      type: Boolean,
      default: false,
    },
  },
};
</script>

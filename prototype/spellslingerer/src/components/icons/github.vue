<template>
  <v-icon>{{ svgPath }}</v-icon>
</template>

<script>
import { mdiGithub } from '@mdi/js';

export default {
  data: () => ({
    svgPath: mdiGithub,
  }),
};
</script>

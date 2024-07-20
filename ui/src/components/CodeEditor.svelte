<script lang="ts">
  import { type LanguageSupport } from "@codemirror/language";

  import CodeMirror from "svelte-codemirror-editor";
  import { nginx } from "@codemirror/legacy-modes/mode/nginx";
  import { StreamLanguage } from "@codemirror/language";
  const nginxLanguage = StreamLanguage.define(
    nginx
  ) as unknown as LanguageSupport;

  let value = ``;
</script>

<CodeMirror bind:value lang={nginxLanguage} />

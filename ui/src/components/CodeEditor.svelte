<script lang="ts">
  import { type LanguageSupport } from "@codemirror/language";

  import CodeMirror from "svelte-codemirror-editor";
  import { nginx } from "@codemirror/legacy-modes/mode/nginx";
  import { StreamLanguage } from "@codemirror/language";
  const nginxLanguage = StreamLanguage.define(
    nginx
  ) as unknown as LanguageSupport;

  import RetroOutline from "./ui/RetroOutline.svelte";

  let value = ``;
</script>

<RetroOutline>
  <CodeMirror bind:value lang={nginxLanguage} />
</RetroOutline>

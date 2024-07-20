<script lang="ts">
	import { getToastState, type ToastType, type ToastMode } from "../lib/toast";
	import XIcon from "lucide-svelte/icons/x";
	import CheckCircleIcon from "lucide-svelte/icons/circle-check";
	import AlertCircleIcon from "lucide-svelte/icons/circle-alert";
	import InfoIcon from "lucide-svelte/icons/info";

	import type XIconType from "lucide-svelte/icons/x";
	import type CheckCircleIconType from "lucide-svelte/icons/check-check";
	import type AlertCircleIconType from "lucide-svelte/icons/circle-alert";
	import type InfoIconType from "lucide-svelte/icons/info";

	import { crossfade } from "svelte/transition";
	import { cubicInOut } from "svelte/easing";
	import RetroOutline from "./ui/RetroOutline.svelte";
  
	export let toast: ToastType;
  
	const toastState = getToastState();
  
	type IconComponent =
	  | typeof XIconType
	  | typeof CheckCircleIconType
	  | typeof AlertCircleIconType
	  | typeof InfoIconType;
  
	const icons: Record<ToastMode, IconComponent> = {
	  success: CheckCircleIcon,
	  error: AlertCircleIcon,
	  info: InfoIcon,
	};
  
	const bgColors: Record<ToastMode, string> = {
	  success: "bg-green-100",
	  error: "bg-red-100",
	  info: "bg-gray-50",
	};
  
	const textColors: Record<ToastMode, string> = {
	  success: "text-green-800",
	  error: "text-red-800",
	  info: "text-gray-900",
	};
  
	$: Icon = icons[toast.type as ToastMode] || InfoIcon;
	$: bgColor = bgColors[toast.type as ToastMode] || "bg-gray-50";
	$: textColor = textColors[toast.type as ToastMode] || "text-gray-800";
  
	const [send, receive] = crossfade({
	  duration: 300,
	  easing: cubicInOut,
	  fallback(node, params) {
		const style = getComputedStyle(node);
		const transform = style.transform === 'none' ? '' : style.transform;
  
		return {
		  duration: 300,
		  easing: cubicInOut,
		  css: t => `
			transform: ${transform} translateY(${(1 - t) * 50}px);
			opacity: ${t}
		  `
		};
	  }
	});
  </script>
  
  <div
	class="relative h-auto w-72 {textColor} px-2 py-1"
	in:receive={{key: toast.id}}
	out:send={{key: toast.id}}
	role="alert"
  >
	<RetroOutline>
	  <div class="flex items-top p-4 {bgColor}">
		<div class="mr-3">
		  <svelte:component this={Icon} class="size-5" />
		</div>
		<div class="flex-1">
		  <h3 class="text-sm font-medium">{toast.title}</h3>
		  {#if toast.message}
			<p class="mt-1 text-xs">{toast.message}</p>
		  {/if}
		</div>
		<div>
		  <button
			class="ml-4 text-gray-600 hover:text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 p-1"
			on:click={() => toastState.remove(toast.id)}
			aria-label="Close"
		  >
			<XIcon class="size-4" />
		  </button>
		</div>
	  </div>
	</RetroOutline>
  </div>
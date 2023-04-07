<script lang="ts">
	import { onMount } from 'svelte';
	import AccentText from '$lib/AccentText.svelte';
	import { dev } from '$app/environment';

	export let provider: string | null;
	let providers: string[] | null = null;

	onMount(async () => {
		if (dev) {
			providers = ['github', 'google', 'twitter'];
			return;
		}
		const res = await fetch('/supported')
		providers = await res.json();
	});
</script>

<p>
	{#if provider}
		Successful authentication with <AccentText text={provider} />! You can now close this window.
	{:else if providers === null}
		Loading...
	{:else if providers.length === 0}
		This server is not configured for any providers.
	{:else if providers.length === 1}
		This server is configured only for <AccentText text={providers[0]} /> provider.
	{:else if providers.length > 1}
		This server is configured for the following providers:
		{#each providers as provider, i}
			<AccentText
				text={provider}
			/>{#if !(providers.length - 2 === i)}{#if !(providers.length === i + 1)},{' '}{:else}.{/if}
			{:else}
				{' '}and{' '}
			{/if}
		{/each}
	{/if}
</p>

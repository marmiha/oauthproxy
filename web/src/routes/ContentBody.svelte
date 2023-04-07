<script lang="ts">
	import AccentText from '$lib/AccentText.svelte';

	export let provider: string | null;
	let providers: string[] = ['google', 'github', 'twitter', 'facebook'];
</script>

<p>
	{#if provider}
		Successful authentication with <AccentText text={provider} />! You can now close this window.
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

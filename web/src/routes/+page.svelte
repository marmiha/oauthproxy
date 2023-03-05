<!-- Right Side Content	-->
<script lang="ts">
	import ContentBody from './ContentBody.svelte';
	import { page } from '$app/stores'
	
    const provider = $page.url.searchParams.get('provider')
</script>

<ContentBody {provider} />

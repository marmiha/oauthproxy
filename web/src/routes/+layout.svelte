<script>
	import GithubIcon from '$lib/GithubLinkIcon.svelte';
	import Logo from '$lib/Logo.svelte';

	const repoUrl = 'https://github.com/gume1a/oauthproxy';
</script>

<div id="header">
	<div class="header-icon" id="github-icon">
		<GithubIcon {repoUrl} />
	</div>
	<div class="header-icon" id="header-logo">
		<Logo max_height="24px" />
	</div>
</div>

<div id="content">
	<div>
		<!-- Left Side Content	-->
		<div id="left-content" class="not-selectable">
			<Logo />
		</div>
		<!-- Right Side Content	-->
		<div id="right-content">
			<h1>oauthproxy</h1>
			<slot />
		</div>
	</div>
</div>

<div class="intro-footer">
	Another project by <a href="github.com/marmiha">marmiha</a> and
	<a href="github.com/gume1a/oauthproxy">friends</a>
</div>

<style>
	:global(html) {
		background: #151619;
		color: #f1f1f1;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif,
			'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol';
		display: flex;
		margin: 0;
		width: 100%;
		height: 100%;
	}

	:global(body) {
		width: 100%;
		margin: 13px;
		display: flex;
	}

	@media (max-width: 433px) {
		#left-content {
			visibility: hidden;
			width: 0;
		}

		#header-logo {
			visibility: visible !important;
		}

		#github-icon {
			right: calc(24px + 9px + 13px) !important;
			top: 13px !important;
		}
	}

	#header {
		position: absolute;
		width: 100%;
	}

	.intro-footer {
		padding-bottom: 1rem;
		font-size: 0.75rem;
		text-align: center;
		position: fixed;
		width: calc(100% - 26px);
		padding: 0 13px 1em 13px;
		bottom: 0;
		left: 0;
	}

	.intro-footer a {
		color: #b44e4e;
	}

	#content {
		margin: auto;
		padding-bottom: 10vh;
	}

	#left-content {
		float: left;
		transition: width 3s ease-in-out;
	}

	#right-content {
		max-width: 230px;
		float: right;
	}

	.header-icon {
		position: fixed;
		right: 13px;
		top: 13px;
	}

	#header-logo {
		position: fixed;
		right: 13px;
		top: 13px;
		visibility: hidden;
	}
	h1 {
		margin-top: 0;
	}

	.not-selectable {
		-webkit-touch-callout: none;
		-webkit-user-select: none;
		-khtml-user-select: none;
		-moz-user-select: none;
		-ms-user-select: none;
		-webkit-user-drag: none;
		user-select: none;
		pointer-events: none;
	}


</style>

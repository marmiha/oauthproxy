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
	<div id="content-wrapper">
		<!-- Left Side Content	-->
		<div id="left-content" class="not-selectable">
			<Logo />
		</div>
		<!-- Right Side Content	-->
		<div id="right-content">
			<h1>oauthproxy</h1>
			<div id="right-content-subheader">
				<slot />
			</div>
		</div>
	</div>
</div>

<div class="intro-footer">
	Another project by <a href="http://github.com/marmiha">marmiha</a> and
	<a href="http://github.com/gume1a/oauthproxy">friends</a>
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

	#content-wrapper {
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
		width: 100%;
		height: 100%;
	}

	#right-content {
		max-width: 230px;
		float: right;
	}

	#left-content {
		float: left;
		transition: width 3s ease-in-out;
	}

	@media (max-width: 433px) {
		#content-wrapper {
			flex-direction: column;
			margin-bottom: 4rem;
			margin-top: 2.5rem;
		}

		#left-content {
			margin-bottom: 0rem;
			display: flex;
			opacity: 0;
			transform: scale(0.7);
			justify-content: center;
			animation: scaleIn 50ms ease-in-out 100ms forwards;
		}

		#right-content > h1 {
			margin-bottom: 0px;
		}

		#right-content {
			margin-top: .5rem;
			text-align: center;
			opacity: 0;
			animation: scaleIn 50ms ease-in-out 125ms forwards !important;
		}

		#right-content-subheader {
			opacity: 0;
			animation: scaleIn 50ms ease-in-out 150ms forwards !important;
		}
	}

	#left-content {
		float: left;
		transition: width 1s ease-out;
	}

	#right-content {
		max-width: 230px;
		float: right;
	}

	@keyframes scaleIn {
		from {
			opacity: 0;
			transform: scale(0.7);
		}

		to {
			opacity: 1;
			transform: scale(1);
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

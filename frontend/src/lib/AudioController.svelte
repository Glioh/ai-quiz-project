<script lang="ts">
    import { onMount, onDestroy } from 'svelte';
    import { faVolumeUp, faVolumeMute } from '@fortawesome/free-solid-svg-icons';
    import Fa from 'svelte-fa';
    import { isMuted } from '../stores/audioStore';

    export let audioFile: string;
    export let iconColor: string = "white";
    export let hoverBgColor: string = "bg-purple-600";
    
    let audio: HTMLAudioElement | null = null;

    onMount(() => {
        if (!audio) {
            audio = new Audio(audioFile);
            audio.loop = true;
            audio.volume = 0.5;
            audio.muted = $isMuted;
            audio.play().catch(error => console.log("Audio autoplay failed:", error));
        }
    });

    onDestroy(() => {
        if (audio) {
            audio.pause();
            audio.currentTime = 0;
            audio = null;
        }
    });

    function toggleMute() {
        $isMuted = !$isMuted;
        if (audio) {
            audio.muted = $isMuted;
        }
    }

    // Only change the audio source if the file changes
    $: if (audio && audioFile && audio.src !== new URL(audioFile, window.location.href).href) {
        audio.src = audioFile;
        audio.play().catch(error => console.log("Audio play failed:", error));
    }
</script>

<button 
    class="p-2 rounded-full transition-colors {`text-${iconColor}`} hover:{hoverBgColor}"
    on:click={toggleMute}
    aria-label={$isMuted ? "Unmute" : "Mute"}
>
    <Fa icon={$isMuted ? faVolumeMute : faVolumeUp} size="1.5x"/>
</button>
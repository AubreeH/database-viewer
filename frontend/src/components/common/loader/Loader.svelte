<div class="loader">
    <div class="loader__inner" />
    <div class="loader__inner" />
    <div class="loader__inner" />
</div>

<style lang="scss">
    @keyframes loader_after {
        0%, 20% {
            margin-inline: 0 calc(100% - .5em);
        }

        80%, 100% {
            margin-inline: calc(100% - .5em) 0;
        }
        
        // 100% {
        //     margin-inline: 0 calc(100% - .5em);
        // }
    }
    
    .loader {
        overflow: hidden;
        width: fit-content;
        border-radius: .5em;
        
        display: grid;
        grid-template-columns: 1fr 6em 1fr;
        grid-template-rows: 1fr .5em 1fr;
        
        &::before {
            grid-column: 2;
            grid-row: 2;

            display: block;
            content: "";

            width: 6em;
            height: .5em;

            background: #f2f2f220;
            border-radius: .25em;
        }

        .loader__inner {
            grid-column: 2;
            grid-row: 2;

            display: block;
            content: "";

            background: #808080;
            animation: loader_after 1.5s infinite alternate;

            &:nth-child(1) {
                animation-timing-function: cubic-bezier(0, 0.25, 1, 0.75);
            }

            &:nth-child(2) {
                animation-timing-function: cubic-bezier(0, 0.6, 1, 1);
            }

            &:nth-child(3) {
                animation-timing-function: cubic-bezier(0, 0, 1, 0.4);
            }

            

            height: .5em;
            width: .5em;
            border-radius: .25em;
        }
    }

    
</style>
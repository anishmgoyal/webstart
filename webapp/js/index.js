(() => {
    const div = document.createElement( "div" );
    div.appendChild( document.createTextNode( "Recommendation: defer scripts and styles if they're not immediately needed to speed up page load" ) );
    div.classList.add( "small-full" );
    
    document.body.appendChild( div );
})();
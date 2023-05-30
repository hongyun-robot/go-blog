const $ = window.$;
// const openOrCloseEl = $('.open-or-close')
// const item = $('.item')
// console.log(openOrCloseEl);
// classify-item
// openOrCloseEl.click(function () {
//     const el = $(this)
//     console.log(el)
//     const isOpen = !el.hasClass('opacity-0')
//     el.toggleClass('opacity-0', isOpen)
//
// })
$('.classify-item').each(function () {
    const that = $(this)
    const btn = that.find('.open-or-close')
    const articleListEl = that.find('ul')
    const articleListElHeight = articleListEl.height()
    articleListEl.height(articleListElHeight)
    btn.click(function () {
        const isOpen = !btn.hasClass('opacity-0')
        btn.toggleClass('opacity-0', isOpen)
        articleListEl.toggleClass('opacity-0', isOpen)
        if (isOpen) {
            articleListEl.height(0)
        } else {
            articleListEl.height(articleListElHeight)
        }
    })
})
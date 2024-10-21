// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "net/url"

func menu(queries url.Values) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"navigation-interaction-area\"></div><nav class=\"navigation navigation-hidden\"><div class=\"navigation--item navigation--control\"><svg class=\"navigation--control--pause\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 320 512\" height=\"32px\" width=\"32px\"><!--!Font Awesome Free 6.6.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2024 Fonticons, Inc.--><path d=\"M48 64C21.5 64 0 85.5 0 112L0 400c0 26.5 21.5 48 48 48l32 0c26.5 0 48-21.5 48-48l0-288c0-26.5-21.5-48-48-48L48 64zm192 0c-26.5 0-48 21.5-48 48l0 288c0 26.5 21.5 48 48 48l32 0c26.5 0 48-21.5 48-48l0-288c0-26.5-21.5-48-48-48l-32 0z\"></path></svg> <svg class=\"navigation--control--play\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 384 512\" height=\"32px\" width=\"32px\"><!--!Font Awesome Free 6.6.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2024 Fonticons, Inc.--><path d=\"M73 39c-14.8-9.1-33.4-9.4-48.5-.9S0 62.6 0 80L0 432c0 17.4 9.4 33.4 24.5 41.9s33.7 8.1 48.5-.9L361 297c14.3-8.7 23-24.2 23-41s-8.7-32.2-23-41L73 39z\"></path></svg></div><div class=\"navigation--item--separator\"></div><div class=\"navigation--item navigation--flush-cache\" hx-get=\"/cache/flush\" hx-swap=\"none\"><svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2 5C2 4.19711 2.43749 3.55194 2.96527 3.08401C3.49422 2.61504 4.20256 2.2384 4.99202 1.94235C6.57833 1.34749 8.70269 1 11 1C13.2973 1 15.4217 1.34749 17.008 1.94235C17.7974 2.2384 18.5058 2.61504 19.0347 3.08401C19.5625 3.55194 20 4.19711 20 5V9.98763C20 10.0333 19.9996 10.0781 19.9989 10.1219C19.9959 10.2876 19.9945 10.3704 19.9431 10.4263C19.8917 10.4822 19.7934 10.4923 19.5967 10.5124C18.9826 10.5753 18.401 10.8669 17.985 11.3266C17.9032 11.4169 17.8623 11.4621 17.8187 11.4792C17.7751 11.4964 17.719 11.4916 17.6067 11.4819C17.4199 11.4659 17.231 11.4577 17.0404 11.4577C15.0435 11.4577 13.2497 12.3522 12.0481 13.7655C11.9557 13.8742 11.9095 13.9285 11.8524 13.956C11.7953 13.9836 11.7298 13.9858 11.599 13.9901C11.3982 13.9967 11.1984 14 11 14C8.6113 14 6.01354 13.5188 4.0508 12.5952C3.64779 12.4055 3.28325 12.2037 2.95806 11.9907C2.15337 11.4637 2 10.9324 2 9.98763V5ZM5.57313 6.13845C4.97883 5.9045 4.55524 5.65279 4.29209 5.41948C3.9354 5.10324 3.9354 4.89676 4.29209 4.58052C4.57279 4.33166 5.03602 4.06185 5.69427 3.81501C7.0034 3.32409 8.87903 3 11 3C13.121 3 14.9966 3.32409 16.3057 3.81501C16.964 4.06185 17.4272 4.33166 17.7079 4.58052C18.0646 4.89676 18.0646 5.10324 17.7079 5.41948C17.4272 5.66834 16.964 5.93815 16.3057 6.18499C14.9966 6.67591 13.121 7 11 7C10.1029 7 9.24969 6.94202 8.46467 6.83796C7.48782 6.70847 6.52272 6.51225 5.57313 6.13845ZM6.21587 10.1237C5.81919 10.0045 5.40095 10.2294 5.2817 10.6261C5.16246 11.0228 5.38736 11.441 5.78404 11.5602C6.42365 11.7525 7.13136 11.9087 7.8874 12.0234C8.29692 12.0856 8.6793 11.804 8.74146 11.3945C8.80362 10.9849 8.52203 10.6026 8.11251 10.5404C7.41722 10.4349 6.77942 10.2932 6.21587 10.1237Z\" fill=\"white\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M10.6635 16.5223C10.7204 16.2736 10.7488 16.1493 10.6912 16.0746C10.6336 15.9999 10.5104 15.9959 10.2641 15.9879C7.84409 15.9091 5.38708 15.4342 3.1992 14.4046C3.00604 14.3137 2.81512 14.2177 2.62747 14.1164C2.3392 13.9608 2.19506 13.883 2.09753 13.9412C2 13.9993 2 14.1584 2 14.4765V18.9998C2 19.8027 2.43749 20.4479 2.96527 20.9158C3.49422 21.3848 4.20256 21.7614 4.99202 22.0575C6.57833 22.6523 8.70269 22.9998 11 22.9998C11.277 22.9998 11.5514 22.9948 11.8227 22.9848C12.219 22.9703 12.4171 22.9631 12.4672 22.8473C12.5174 22.7315 12.3782 22.5719 12.1 22.2526C11.1041 21.1097 10.5 19.6168 10.5 17.9789C10.5 17.4778 10.5565 16.9903 10.6635 16.5223ZM6.21587 17.1237C5.81919 17.0045 5.40095 17.2294 5.2817 17.6261C5.16246 18.0228 5.38736 18.441 5.78404 18.5602C6.42365 18.7525 7.13136 18.9087 7.8874 19.0234C8.29692 19.0856 8.6793 18.804 8.74146 18.3945C8.80362 17.9849 8.52203 17.6026 8.11251 17.5404C7.41722 17.4349 6.77942 17.2932 6.21587 17.1237Z\" fill=\"white\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M19.6056 12.0284C19.0689 12.159 18.7398 12.6999 18.8704 13.2365L18.8873 13.3058C18.3149 13.0811 17.6917 12.9576 17.0404 12.9576C14.2612 12.9576 12 15.2012 12 17.9788C12 20.7564 14.2612 23 17.0404 23C19.4765 23 21.5117 21.278 21.9798 18.9829C22.0902 18.4417 21.741 17.9136 21.1998 17.8032C20.6587 17.6928 20.1305 18.0421 20.0202 18.5832C19.7396 19.959 18.5137 21 17.0404 21C15.3567 21 14 19.6429 14 17.9788C14 16.3147 15.3567 14.9576 17.0404 14.9576C17.7271 14.9576 18.3577 15.1828 18.8659 15.5627C18.8901 15.5808 18.9151 15.5978 18.9408 15.6136L19.8888 16.1967C20.2341 16.4091 20.6734 16.392 21.0011 16.1535C21.3288 15.915 21.4802 15.5023 21.3844 15.1084L20.8137 12.7635C20.6831 12.2269 20.1422 11.8978 19.6056 12.0284Z\" fill=\"white\"></path></svg></div><div class=\"navigation--item--separator navigation--fullscreen-separator\"></div><div class=\"navigation--item navigation--fullscreen\"><svg class=\"navigation--fullscreen--enter\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\" height=\"32px\" width=\"32px\"><!--!Font Awesome Free 6.6.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2024 Fonticons, Inc.--><path d=\"M344 0L488 0c13.3 0 24 10.7 24 24l0 144c0 9.7-5.8 18.5-14.8 22.2s-19.3 1.7-26.2-5.2l-39-39-87 87c-9.4 9.4-24.6 9.4-33.9 0l-32-32c-9.4-9.4-9.4-24.6 0-33.9l87-87L327 41c-6.9-6.9-8.9-17.2-5.2-26.2S334.3 0 344 0zM168 512L24 512c-13.3 0-24-10.7-24-24L0 344c0-9.7 5.8-18.5 14.8-22.2s19.3-1.7 26.2 5.2l39 39 87-87c9.4-9.4 24.6-9.4 33.9 0l32 32c9.4 9.4 9.4 24.6 0 33.9l-87 87 39 39c6.9 6.9 8.9 17.2 5.2 26.2s-12.5 14.8-22.2 14.8z\"></path></svg> <svg class=\"navigation--fullscreen--exit\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\" height=\"32px\" width=\"32px\"><!--!Font Awesome Free 6.6.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2024 Fonticons, Inc.--><path d=\"M439 7c9.4-9.4 24.6-9.4 33.9 0l32 32c9.4 9.4 9.4 24.6 0 33.9l-87 87 39 39c6.9 6.9 8.9 17.2 5.2 26.2s-12.5 14.8-22.2 14.8l-144 0c-13.3 0-24-10.7-24-24l0-144c0-9.7 5.8-18.5 14.8-22.2s19.3-1.7 26.2 5.2l39 39L439 7zM72 272l144 0c13.3 0 24 10.7 24 24l0 144c0 9.7-5.8 18.5-14.8 22.2s-19.3 1.7-26.2-5.2l-39-39L73 505c-9.4 9.4-24.6 9.4-33.9 0L7 473c-9.4-9.4-9.4-24.6 0-33.9l87-87L55 313c-6.9-6.9-8.9-17.2-5.2-26.2s12.5-14.8 22.2-14.8z\"></path></svg></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate

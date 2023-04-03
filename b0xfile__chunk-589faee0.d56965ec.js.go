package knife4go
import (
	"os"
)

func init() {

	f, err := FS.OpenFile(CTX, "/chunk-589faee0.d56965ec.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([] byte(base64Decoding("KHdpbmRvdy53ZWJwYWNrSnNvbnA9d2luZG93LndlYnBhY2tKc29ucHx8W10pLnB1c2goW1siY2h1bmstNTg5ZmFlZTAiXSx7IjNjNjAiOmZ1bmN0aW9uKGUsdCxuKXsidXNlIHN0cmljdCI7bi5yKHQpO3ZhciBvPShuKCJkM2I3IiksbigiM2NhMyIpLG4oImRkYjAiKSxuKCJiNjgwIiksbigiMTU5YiIpLG4oImFjMWYiKSxuKCI1MzE5IiksbigiMjVmMCIpLG4oIjE0ZDkiKSxuKCJhMTViIiksbigiYjFjNyIpKSxyPW4oImIzMTEiKSxzPW4ubihyKSxpPXtwcm9wczp7YXBpOnt0eXBlOk9iamVjdCxyZXF1aXJlZDohMH0sc3dhZ2dlckluc3RhbmNlOnt0eXBlOk9iamVjdCxyZXF1aXJlZDohMH0sZGVidWdTZW5kOnt0eXBlOkJvb2xlYW4sZGVmYXVsdDohMX0scmVzcG9uc2VIZWFkZXJzOnt0eXBlOkFycmF5fSxyZXNwb25zZVJhd1RleHQ6e3R5cGU6U3RyaW5nLGRlZmF1bHQ6IiJ9LHJlc3BvbnNlQ3VybFRleHQ6e3R5cGU6U3RyaW5nLGRlZmF1bHQ6IiJ9LHJlc3BvbnNlU3RhdHVzOnt0eXBlOk9iamVjdH0scmVzcG9uc2VDb250ZW50Ont0eXBlOk9iamVjdH0scmVzcG9uc2VGaWVsZERlc2NyaXB0aW9uQ2hlY2tlZDp7dHlwZTpCb29sZWFuLGRlZmF1bHQ6ITB9fSxjb21wb25lbnRzOntFZGl0b3JEZWJ1Z1Nob3c6ZnVuY3Rpb24oKXtyZXR1cm4gUHJvbWlzZS5hbGwoW24uZSgiY2h1bmstM2I4ODhhNjUiKSxuLmUoImNodW5rLTBmZDY3NzE2Iiksbi5lKCJjaHVuay0yMTQyMThmMCIpLG4uZSgiY2h1bmstNzM1YzY3NWMiKV0pLnRoZW4obi5iaW5kKG51bGwsImIzZWUiKSl9fSxkYXRhOmZ1bmN0aW9uKCl7cmV0dXJue3BhZ2luYXRpb246ITEsaTE4bjpudWxsLGJhc2U2NEltYWdlOiExLGRlYnVnUmVzcG9uc2U6ITAscmVzcG9uc2VIZWFkZXJDb2x1bW46W119fSx3YXRjaDp7bGFuZ3VhZ2U6ZnVuY3Rpb24oZSx0KXt0aGlzLmluaXRJMThuKCl9fSxjb21wdXRlZDp7bGFuZ3VhZ2U6ZnVuY3Rpb24oKXtyZXR1cm4gdGhpcy4kc3RvcmUuc3RhdGUuZ2xvYmFscy5sYW5ndWFnZX0scmVzcG9uc2VTaXplVGV4dDpmdW5jdGlvbigpe3ZhciBlPSIwIEIiLHQ9dGhpcy5yZXNwb25zZVN0YXR1cztpZihudWxsIT10JiZudWxsIT10KXt2YXIgbj10LnNpemUsbz0obi8xMDI0KS50b0ZpeGVkKDIpLHI9KG4vMTAyNC8xMDI0KS50b0ZpeGVkKDIpO2U9bz4xP28rIiBLQiI6cj4xP3IrIiBNQiI6bisiIEIifXJldHVybiBlfX0sY3JlYXRlZDpmdW5jdGlvbigpe3RoaXMuaW5pdEkxOG4oKSx0aGlzLmNvcHlSYXdUZXh0KCksdGhpcy5jb3B5Q3VybFRleHQoKX0sbWV0aG9kczp7Z2V0Q3VycmVudEkxOG5JbnN0YW5jZTpmdW5jdGlvbigpe3JldHVybiB0aGlzLiRpMThuLm1lc3NhZ2VzW3RoaXMubGFuZ3VhZ2VdfSxiYXNlNjRJbml0OmZ1bmN0aW9uKCl7dmFyIGU9by5hLmdldFZhbHVlKHRoaXMucmVzcG9uc2VDb250ZW50LCJiYXNlNjQiLCIiLCEwKTtvLmEuc3RyTm90QmxhbmsoZSkmJih0aGlzLmJhc2U2NEltYWdlPSEwKX0saW5pdEkxOG46ZnVuY3Rpb24oKXt0aGlzLmkxOG49dGhpcy5nZXRDdXJyZW50STE4bkluc3RhbmNlKCksdGhpcy5yZXNwb25zZUhlYWRlckNvbHVtbj10aGlzLmkxOG4udGFibGUuZGVidWdSZXNwb25zZUhlYWRlckNvbHVtbnN9LGNvcHlSYXdUZXh0OmZ1bmN0aW9uKCl7dmFyIGU9dGhpcyx0PSJidG5EZWJ1Z0NvcHlSYXciK3RoaXMuYXBpLmlkLG49bmV3IHMuYSgiIyIrdCx7dGV4dDpmdW5jdGlvbigpe3JldHVybiBlLnJlc3BvbnNlUmF3VGV4dH19KSxvPXRoaXMuaTE4bi5tZXNzYWdlLmNvcHkucmF3LnN1Y2Nlc3Mscj10aGlzLmkxOG4ubWVzc2FnZS5jb3B5LnJhdy5mYWlsO24ub24oInN1Y2Nlc3MiLChmdW5jdGlvbih0KXtlLiRtZXNzYWdlLmluZm8obyl9KSksbi5vbigiZXJyb3IiLChmdW5jdGlvbih0KXtlLiRtZXNzYWdlLmluZm8ocil9KSl9LGNvcHlDdXJsVGV4dDpmdW5jdGlvbigpe3ZhciBlPXRoaXMsdD0iYnRuRGVidWdDb3B5Q3VybCIrdGhpcy5hcGkuaWQsbj1uZXcgcy5hKCIjIit0LHt0ZXh0OmZ1bmN0aW9uKCl7cmV0dXJuIGUucmVzcG9uc2VDdXJsVGV4dH19KSxvPXRoaXMuaTE4bi5tZXNzYWdlLmNvcHkuY3VybC5zdWNjZXNzLHI9dGhpcy5pMThuLm1lc3NhZ2UuY29weS5jdXJsLmZhaWw7bi5vbigic3VjY2VzcyIsKGZ1bmN0aW9uKHQpe2UuJG1lc3NhZ2UuaW5mbyhvKX0pKSxuLm9uKCJlcnJvciIsKGZ1bmN0aW9uKHQpe2UuJG1lc3NhZ2UuaW5mbyhyKX0pKX0scmVzZXRSZXNwb25zZUNvbnRlbnQ6ZnVuY3Rpb24oKXtpZihudWxsIT10aGlzLnJlc3BvbnNlQ29udGVudCYmImpzb24iPT10aGlzLnJlc3BvbnNlQ29udGVudC5tb2RlKXt2YXIgZT10aGlzLnJlc3BvbnNlQ29udGVudC50ZXh0O3RoaXMucmVzcG9uc2VDb250ZW50LnRleHQ9by5hLmpzb241c3RyaW5naWZ5KG8uYS5qc29uNXBhcnNlKGUpKX19LHNob3dGaWVsZERlc0NoYW5nZTpmdW5jdGlvbihlKXt2YXIgdD1lLnRhcmdldC5jaGVja2VkO3RoaXMuJGVtaXQoImRlYnVnU2hvd0ZpZWxkRGVzY3JpcHRpb25DaGFuZ2UiLHQpLHRoaXMudG9nZ2xlRmllbGREZXNjcmlwdGlvbih0KX0sZGVidWdFZGl0b3JDaGFuZ2U6ZnVuY3Rpb24oZSl7dGhpcy4kZW1pdCgiZGVidWdFZGl0b3JDaGFuZ2UiLGUpfSx0b2dnbGVGaWVsZERlc2NyaXB0aW9uOmZ1bmN0aW9uKGUpe3ZhciB0PSJyZXNwb25zZUVkaXRvckNvbnRlbnQiK3RoaXMuYXBpLmlkLG49ZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQodCkuZ2V0RWxlbWVudHNCeUNsYXNzTmFtZSgia25pZmU0ai1kZWJ1Zy1lZGl0b3ItZmllbGQtZGVzY3JpcHRpb24iKTtvLmEuYXJyTm90RW1wdHkobik/bi5mb3JFYWNoKChmdW5jdGlvbih0KXt0LnN0eWxlLmRpc3BsYXk9ZT8iYmxvY2siOiJub25lIn0pKTp0aGlzLnNob3dFZGl0b3JGaWVsZEFueVdheSgpfSxzaG93RWRpdG9yRmllbGREZXNjcmlwdGlvbjpmdW5jdGlvbihlKXt2YXIgdD10aGlzO28uYS5jaGVja1VuZGVmaW5lZChlKSYmcGFyc2VJbnQoZSk8PTIwMCYmc2V0VGltZW91dCgoZnVuY3Rpb24oKXt0LnNob3dFZGl0b3JGaWVsZFdhaXQoKX0pLDEwMCl9LHNob3dFZGl0b3JGaWVsZFdhaXQ6ZnVuY3Rpb24oKXt0aGlzLmRlYnVnU2VuZCYmdGhpcy5yZXNwb25zZUZpZWxkRGVzY3JpcHRpb25DaGVja2VkJiYianNvbiI9PXRoaXMucmVzcG9uc2VDb250ZW50Lm1vZGUmJnRoaXMuc2hvd0VkaXRvckZpZWxkQW55V2F5KCl9LHNob3dFZGl0b3JGaWVsZEFueVdheTpmdW5jdGlvbigpe3ZhciBlPXRoaXMuc3dhZ2dlckluc3RhbmNlLHQ9dGhpcy5hcGkuZ2V0SHR0cFN1Y2Nlc3NDb2RlT2JqZWN0KCksbj0icmVzcG9uc2VFZGl0b3JDb250ZW50Iit0aGlzLmFwaS5pZCxyPWRvY3VtZW50LmdldEVsZW1lbnRCeUlkKG4pLHM9W10saT1yLmdldEVsZW1lbnRzQnlDbGFzc05hbWUoImFjZV90ZXh0LWxheWVyIiksYT0wLHU9ci5xdWVyeVNlbGVjdG9yKCIuYWNlX3ByaW50LW1hcmdpbiIpO2lmKG8uYS5jaGVja1VuZGVmaW5lZCh1KSYmby5hLmNoZWNrVW5kZWZpbmVkKHUuc3R5bGUpJiYoYT11LnN0eWxlLmxlZnQpLGkubGVuZ3RoPjApZm9yKHZhciBjPWlbMF0uZ2V0RWxlbWVudHNCeUNsYXNzTmFtZSgiYWNlX2xpbmUiKSxsPTA7bDxjLmxlbmd0aDtsKyspe3ZhciBwPWNbbF0sZD1wLmdldEVsZW1lbnRzQnlDbGFzc05hbWUoImFjZV92YXJpYWJsZSIpLGY9bnVsbDtpZihvLmEuYXJyTm90RW1wdHkoZCkpe2Y9by5hLnRvU3RyaW5nKGRbMF0uaW5uZXJIVE1MLCIiKS5yZXBsYWNlKC9eIiguKikiJC9nLCIkMSIpO3ZhciB5PXAuZ2V0RWxlbWVudHNCeUNsYXNzTmFtZSgia25pZmU0ai1kZWJ1Zy1lZGl0b3ItZmllbGQtZGVzY3JpcHRpb24iKTtpZighby5hLmFyck5vdEVtcHR5KHkpKXt2YXIgZz1kb2N1bWVudC5jcmVhdGVFbGVtZW50KCJzcGFuIik7Zy5jbGFzc05hbWU9ImtuaWZlNGotZGVidWctZWRpdG9yLWZpZWxkLWRlc2NyaXB0aW9uIixnLmlubmVySFRNTD10LnJlc3BvbnNlRGVzY3JpcHRpb25GaW5kKHMsZixlKSxnLnN0eWxlLmxlZnQ9YSxwLmFwcGVuZENoaWxkKGcpfX12YXIgaD1wLmdldEVsZW1lbnRzQnlDbGFzc05hbWUoImFjZV9wYXJlbiIpO2lmKG8uYS5hcnJOb3RFbXB0eShoKSl7Zm9yKHZhciBiPVtdLG09MDttPGgubGVuZ3RoO20rKyliLnB1c2goaFttXS5pbm5lckhUTUwpO3N3aXRjaChiLmpvaW4oIiIpKXtjYXNlIlsiOmNhc2UieyI6cy5wdXNoKGZ8fDApO2JyZWFrO2Nhc2UifSI6Y2FzZSJdIjpzLnBvcCgpfX19fX19LGE9bigiMjg3NyIpLHU9T2JqZWN0KGEuYSkoaSwoZnVuY3Rpb24oKXt2YXIgZT10aGlzLHQ9ZS5fc2VsZi5fYztyZXR1cm4gdCgiYS1yb3ciLHtzdGF0aWNDbGFzczoia25pZmU0ai1kZWJ1Zy1yZXNwb25zZSJ9LFtlLmRlYnVnU2VuZD90KCJhLXJvdyIsW3QoImEtdGFicyIse2F0dHJzOntkZWZhdWx0QWN0aXZlS2V5OiJkZWJ1Z1Jlc3BvbnNlIn19LFt0KCJ0ZW1wbGF0ZSIse3Nsb3Q6InRhYkJhckV4dHJhQ29udGVudCJ9LFtlLnJlc3BvbnNlU3RhdHVzP3QoImEtcm93Iix7c3RhdGljQ2xhc3M6ImtuaWZlNGotZGVidWctc3RhdHVzIn0sW3QoInNwYW4iLFt0KCJhLWNoZWNrYm94Iix7YXR0cnM6e2RlZmF1bHRDaGVja2VkOmUucmVzcG9uc2VGaWVsZERlc2NyaXB0aW9uQ2hlY2tlZH0sb246e2NoYW5nZTplLnNob3dGaWVsZERlc0NoYW5nZX19LFt0KCJzcGFuIix7c3RhdGljU3R5bGU6e2NvbG9yOiIjOTE5MTkxIn0sZG9tUHJvcHM6e2lubmVySFRNTDplLl9zKGUuJHQoImRlYnVnLnJlc3BvbnNlLnNob3dEZXMiKSl9fSxbZS5fdigi5pi+56S66K+05piOIildKV0pXSwxKSx0KCJzcGFuIix7c3RhdGljQ2xhc3M6ImtleSIsZG9tUHJvcHM6e2lubmVySFRNTDplLl9zKGUuJHQoImRlYnVnLnJlc3BvbnNlLmNvZGUiKSl9fSxbZS5fdigi5ZON5bqU56CBOiIpXSksdCgic3BhbiIse3N0YXRpY0NsYXNzOiJ2YWx1ZSJ9LFtlLl92KGUuX3MoZS5yZXNwb25zZVN0YXR1cy5jb2RlKSldKSx0KCJzcGFuIix7c3RhdGljQ2xhc3M6ImtleSIsZG9tUHJvcHM6e2lubmVySFRNTDplLl9zKGUuJHQoImRlYnVnLnJlc3BvbnNlLmNvc3QiKSl9fSxbZS5fdigi6ICX5pe2OiIpXSksdCgic3BhbiIse3N0YXRpY0NsYXNzOiJ2YWx1ZSJ9LFtlLl92KGUuX3MoZS5yZXNwb25zZVN0YXR1cy5jb3N0KSldKSx0KCJzcGFuIix7c3RhdGljQ2xhc3M6ImtleSIsZG9tUHJvcHM6e2lubmVySFRNTDplLl9zKGUuJHQoImRlYnVnLnJlc3BvbnNlLnNpemUiKSl9fSxbZS5fdigi5aSn5bCPOiIpXSksdCgic3BhbiIse3N0YXRpY0NsYXNzOiJ2YWx1ZSJ9LFtlLl92KGUuX3MoZS5yZXNwb25zZVNpemVUZXh0KSsiICIpXSldKTplLl9lKCldLDEpLHQoImEtdGFiLXBhbmUiLHtrZXk6ImRlYnVnUmVzcG9uc2UiLGF0dHJzOnt0YWI6ZS5pMThuLmRlYnVnLnJlc3BvbnNlLmNvbnRlbnR9fSxbZS5yZXNwb25zZUNvbnRlbnQ/dCgiYS1yb3ciLFtlLnJlc3BvbnNlQ29udGVudC5ibG9iRmxhZz90KCJhLXJvdyIsW2UucmVzcG9uc2VDb250ZW50LmltYWdlRmxhZz90KCJkaXYiLFt0KCJpbWciLHthdHRyczp7c3JjOmUucmVzcG9uc2VDb250ZW50LmJsb2JVcmx9fSldKTp0KCJkaXYiLFt0KCJhLWJ1dHRvbiIse2F0dHJzOnt0eXBlOiJsaW5rIixocmVmOmUucmVzcG9uc2VDb250ZW50LmJsb2JVcmwsZG93bmxvYWQ6ZS5yZXNwb25zZUNvbnRlbnQuYmxvYkZpbGVOYW1lfSxkb21Qcm9wczp7aW5uZXJIVE1MOmUuX3MoZS4kdCgiZGVidWcucmVzcG9uc2UuZG93bmxvYWQiKSl9fSxbZS5fdigi5LiL6L295paH5Lu2IildKV0sMSldKTp0KCJhLXJvdyIse2F0dHJzOntpZDoicmVzcG9uc2VFZGl0b3JDb250ZW50IitlLmFwaS5pZH19LFt0KCJlZGl0b3ItZGVidWctc2hvdyIse2F0dHJzOntkZWJ1Z1Jlc3BvbnNlOmUuZGVidWdSZXNwb25zZSx2YWx1ZTplLnJlc3BvbnNlQ29udGVudC50ZXh0LG1vZGU6ZS5yZXNwb25zZUNvbnRlbnQubW9kZX0sb246e3Nob3dEZXNjcmlwdGlvbjplLnNob3dFZGl0b3JGaWVsZERlc2NyaXB0aW9uLGRlYnVnRWRpdG9yQ2hhbmdlOmUuZGVidWdFZGl0b3JDaGFuZ2V9fSldLDEpXSwxKTplLl9lKCldLDEpLHQoImEtdGFiLXBhbmUiLHtrZXk6ImRlYnVnUmF3IixhdHRyczp7dGFiOiJSYXciLGZvcmNlUmVuZGVyOiIifX0sW3QoImEtcm93Iix7c3RhdGljQ2xhc3M6ImtuaWZlNGotZGVidWctcmVzcG9uc2UtbXQifSxbdCgiYS1idXR0b24iLHthdHRyczp7aWQ6ImJ0bkRlYnVnQ29weVJhdyIrZS5hcGkuaWQsdHlwZToicHJpbWFyeSJ9fSxbdCgiYS1pY29uIix7YXR0cnM6e3R5cGU6ImNvcHkifX0pLGUuX3YoIiAiKSx0KCJzcGFuIix7ZG9tUHJvcHM6e2lubmVySFRNTDplLl9zKGUuJHQoImRlYnVnLnJlc3BvbnNlLmNvcHkiKSl9fSxbZS5fdigi5aSN5Yi2IildKV0sMSldLDEpLHQoImEtcm93Iix7c3RhdGljQ2xhc3M6ImtuaWZlNGotZGVidWctcmVzcG9uc2UtbXQifSxbdCgiYS10ZXh0YXJlYSIse2F0dHJzOntyb3dzOjEwLHZhbHVlOmUucmVzcG9uc2VSYXdUZXh0fX0pXSwxKV0sMSksdCgiYS10YWItcGFuZSIse2tleToiZGVidWdIZWFkZXJzIixhdHRyczp7dGFiOiJIZWFkZXJzIn19LFt0KCJhLXJvdyIse3N0YXRpY0NsYXNzOiJrbmlmZTRqLWRlYnVnLXJlc3BvbnNlLW10In0sW3QoImEtdGFibGUiLHthdHRyczp7Ym9yZGVyZWQ6IiIsc2l6ZToic21hbGwiLGNvbHVtbnM6ZS5yZXNwb25zZUhlYWRlckNvbHVtbixwYWdpbmF0aW9uOmUucGFnaW5hdGlvbixkYXRhU291cmNlOmUucmVzcG9uc2VIZWFkZXJzLHJvd0tleToiaWQifX0pXSwxKV0sMSksdCgiYS10YWItcGFuZSIse2tleToiZGVidWdDdXJsIixhdHRyczp7dGFiOiJDdXJsIn19LFt0KCJhLXJvdyIse3N0YXRpY0NsYXNzOiJrbmlmZTRqLWRlYnVnLXJlc3BvbnNlLW10In0sW3QoImEtYnV0dG9uIix7YXR0cnM6e2lkOiJidG5EZWJ1Z0NvcHlDdXJsIitlLmFwaS5pZCx0eXBlOiJwcmltYXJ5In19LFt0KCJhLWljb24iLHthdHRyczp7dHlwZToiY29weSJ9fSksZS5fdigiICIpLHQoInNwYW4iLHtkb21Qcm9wczp7aW5uZXJIVE1MOmUuX3MoZS4kdCgiZGVidWcucmVzcG9uc2UuY29weSIpKX19LFtlLl92KCLlpI3liLYiKV0pXSwxKV0sMSksdCgiYS1yb3ciLHtzdGF0aWNDbGFzczoia25pZmU0ai1kZWJ1Zy1yZXNwb25zZS1tdCJ9LFt0KCJwcmUiLHtzdGF0aWNDbGFzczoia25pZmU0ai1kZWJ1Zy1yZXNwb25zZS1jdXJsIn0sW2UuX3YoZS5fcyhlLnJlc3BvbnNlQ3VybFRleHQpKV0pXSldLDEpLG51bGwhPWUucmVzcG9uc2VDb250ZW50JiZudWxsIT1lLnJlc3BvbnNlQ29udGVudC5iYXNlNjQmJiIiIT1lLnJlc3BvbnNlQ29udGVudC5iYXNlNjQ/dCgiYS10YWItcGFuZSIse2tleToiZGVidWdCYXNlNjRJbWciLGF0dHJzOnt0YWI6IkJhc2U2NEltZyJ9fSxbdCgiYS1yb3ciLHtzdGF0aWNDbGFzczoia25pZmU0ai1kZWJ1Zy1yZXNwb25zZS1tdCJ9LFt0KCJpbWciLHthdHRyczp7c3JjOmUucmVzcG9uc2VDb250ZW50LmJhc2U2NH19KV0pXSwxKTplLl9lKCldLDIpXSwxKTp0KCJhLXJvdyIpXSwxKX0pLFtdLCExLG51bGwsbnVsbCxudWxsKTt0LmRlZmF1bHQ9dS5leHBvcnRzfSxiMzExOmZ1bmN0aW9uKGUsdCxuKXsKLyohCiAqIGNsaXBib2FyZC5qcyB2Mi4wLjExCiAqIGh0dHBzOi8vY2xpcGJvYXJkanMuY29tLwogKgogKiBMaWNlbnNlZCBNSVQgwqkgWmVubyBSb2NoYQogKi8KZS5leHBvcnRzPWZ1bmN0aW9uKCl7dmFyIGU9ezY4NjpmdW5jdGlvbihlLHQsbil7InVzZSBzdHJpY3QiO24uZCh0LHtkZWZhdWx0OmZ1bmN0aW9uKCl7cmV0dXJuIEV9fSk7dmFyIG89bigyNzkpLHI9bi5uKG8pLHM9bigzNzApLGk9bi5uKHMpLGE9big4MTcpLHU9bi5uKGEpO2Z1bmN0aW9uIGMoZSl7dHJ5e3JldHVybiBkb2N1bWVudC5leGVjQ29tbWFuZChlKX1jYXRjaChlKXtyZXR1cm4hMX19dmFyIGw9ZnVuY3Rpb24oZSl7dmFyIHQ9dSgpKGUpO3JldHVybiBjKCJjdXQiKSx0fSxwPWZ1bmN0aW9uKGUsdCl7dmFyIG49ZnVuY3Rpb24oZSl7dmFyIHQ9InJ0bCI9PT1kb2N1bWVudC5kb2N1bWVudEVsZW1lbnQuZ2V0QXR0cmlidXRlKCJkaXIiKSxuPWRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoInRleHRhcmVhIik7bi5zdHlsZS5mb250U2l6ZT0iMTJwdCIsbi5zdHlsZS5ib3JkZXI9IjAiLG4uc3R5bGUucGFkZGluZz0iMCIsbi5zdHlsZS5tYXJnaW49IjAiLG4uc3R5bGUucG9zaXRpb249ImFic29sdXRlIixuLnN0eWxlW3Q/InJpZ2h0IjoibGVmdCJdPSItOTk5OXB4Ijt2YXIgbz13aW5kb3cucGFnZVlPZmZzZXR8fGRvY3VtZW50LmRvY3VtZW50RWxlbWVudC5zY3JvbGxUb3A7cmV0dXJuIG4uc3R5bGUudG9wPSIiLmNvbmNhdChvLCJweCIpLG4uc2V0QXR0cmlidXRlKCJyZWFkb25seSIsIiIpLG4udmFsdWU9ZSxufShlKTt0LmNvbnRhaW5lci5hcHBlbmRDaGlsZChuKTt2YXIgbz11KCkobik7cmV0dXJuIGMoImNvcHkiKSxuLnJlbW92ZSgpLG99LGQ9ZnVuY3Rpb24oZSl7dmFyIHQ9YXJndW1lbnRzLmxlbmd0aD4xJiZ2b2lkIDAhPT1hcmd1bWVudHNbMV0/YXJndW1lbnRzWzFdOntjb250YWluZXI6ZG9jdW1lbnQuYm9keX0sbj0iIjtyZXR1cm4ic3RyaW5nIj09dHlwZW9mIGU/bj1wKGUsdCk6ZSBpbnN0YW5jZW9mIEhUTUxJbnB1dEVsZW1lbnQmJiFbInRleHQiLCJzZWFyY2giLCJ1cmwiLCJ0ZWwiLCJwYXNzd29yZCJdLmluY2x1ZGVzKG51bGw9PWU/dm9pZCAwOmUudHlwZSk/bj1wKGUudmFsdWUsdCk6KG49dSgpKGUpLGMoImNvcHkiKSksbn07ZnVuY3Rpb24gZihlKXtyZXR1cm4oZj0iZnVuY3Rpb24iPT10eXBlb2YgU3ltYm9sJiYic3ltYm9sIj09dHlwZW9mIFN5bWJvbC5pdGVyYXRvcj9mdW5jdGlvbihlKXtyZXR1cm4gdHlwZW9mIGV9OmZ1bmN0aW9uKGUpe3JldHVybiBlJiYiZnVuY3Rpb24iPT10eXBlb2YgU3ltYm9sJiZlLmNvbnN0cnVjdG9yPT09U3ltYm9sJiZlIT09U3ltYm9sLnByb3RvdHlwZT8ic3ltYm9sIjp0eXBlb2YgZX0pKGUpfXZhciB5PWZ1bmN0aW9uKCl7dmFyIGU9YXJndW1lbnRzLmxlbmd0aD4wJiZ2b2lkIDAhPT1hcmd1bWVudHNbMF0/YXJndW1lbnRzWzBdOnt9LHQ9ZS5hY3Rpb24sbj12b2lkIDA9PT10PyJjb3B5Ijp0LG89ZS5jb250YWluZXIscj1lLnRhcmdldCxzPWUudGV4dDtpZigiY29weSIhPT1uJiYiY3V0IiE9PW4pdGhyb3cgbmV3IEVycm9yKCdJbnZhbGlkICJhY3Rpb24iIHZhbHVlLCB1c2UgZWl0aGVyICJjb3B5IiBvciAiY3V0IicpO2lmKHZvaWQgMCE9PXIpe2lmKCFyfHwib2JqZWN0IiE9PWYocil8fDEhPT1yLm5vZGVUeXBlKXRocm93IG5ldyBFcnJvcignSW52YWxpZCAidGFyZ2V0IiB2YWx1ZSwgdXNlIGEgdmFsaWQgRWxlbWVudCcpO2lmKCJjb3B5Ij09PW4mJnIuaGFzQXR0cmlidXRlKCJkaXNhYmxlZCIpKXRocm93IG5ldyBFcnJvcignSW52YWxpZCAidGFyZ2V0IiBhdHRyaWJ1dGUuIFBsZWFzZSB1c2UgInJlYWRvbmx5IiBpbnN0ZWFkIG9mICJkaXNhYmxlZCIgYXR0cmlidXRlJyk7aWYoImN1dCI9PT1uJiYoci5oYXNBdHRyaWJ1dGUoInJlYWRvbmx5Iil8fHIuaGFzQXR0cmlidXRlKCJkaXNhYmxlZCIpKSl0aHJvdyBuZXcgRXJyb3IoJ0ludmFsaWQgInRhcmdldCIgYXR0cmlidXRlLiBZb3UgY2FuXCd0IGN1dCB0ZXh0IGZyb20gZWxlbWVudHMgd2l0aCAicmVhZG9ubHkiIG9yICJkaXNhYmxlZCIgYXR0cmlidXRlcycpfXJldHVybiBzP2Qocyx7Y29udGFpbmVyOm99KTpyPyJjdXQiPT09bj9sKHIpOmQocix7Y29udGFpbmVyOm99KTp2b2lkIDB9O2Z1bmN0aW9uIGcoZSl7cmV0dXJuKGc9ImZ1bmN0aW9uIj09dHlwZW9mIFN5bWJvbCYmInN5bWJvbCI9PXR5cGVvZiBTeW1ib2wuaXRlcmF0b3I/ZnVuY3Rpb24oZSl7cmV0dXJuIHR5cGVvZiBlfTpmdW5jdGlvbihlKXtyZXR1cm4gZSYmImZ1bmN0aW9uIj09dHlwZW9mIFN5bWJvbCYmZS5jb25zdHJ1Y3Rvcj09PVN5bWJvbCYmZSE9PVN5bWJvbC5wcm90b3R5cGU/InN5bWJvbCI6dHlwZW9mIGV9KShlKX1mdW5jdGlvbiBoKGUsdCl7Zm9yKHZhciBuPTA7bjx0Lmxlbmd0aDtuKyspe3ZhciBvPXRbbl07by5lbnVtZXJhYmxlPW8uZW51bWVyYWJsZXx8ITEsby5jb25maWd1cmFibGU9ITAsInZhbHVlImluIG8mJihvLndyaXRhYmxlPSEwKSxPYmplY3QuZGVmaW5lUHJvcGVydHkoZSxvLmtleSxvKX19ZnVuY3Rpb24gYihlLHQpe3JldHVybihiPU9iamVjdC5zZXRQcm90b3R5cGVPZnx8ZnVuY3Rpb24oZSx0KXtyZXR1cm4gZS5fX3Byb3RvX189dCxlfSkoZSx0KX1mdW5jdGlvbiBtKGUpe3ZhciB0PWZ1bmN0aW9uKCl7aWYoInVuZGVmaW5lZCI9PXR5cGVvZiBSZWZsZWN0fHwhUmVmbGVjdC5jb25zdHJ1Y3QpcmV0dXJuITE7aWYoUmVmbGVjdC5jb25zdHJ1Y3Quc2hhbSlyZXR1cm4hMTtpZigiZnVuY3Rpb24iPT10eXBlb2YgUHJveHkpcmV0dXJuITA7dHJ5e3JldHVybiBEYXRlLnByb3RvdHlwZS50b1N0cmluZy5jYWxsKFJlZmxlY3QuY29uc3RydWN0KERhdGUsW10sKGZ1bmN0aW9uKCl7fSkpKSwhMH1jYXRjaChlKXtyZXR1cm4hMX19KCk7cmV0dXJuIGZ1bmN0aW9uKCl7dmFyIG4sbz1DKGUpO2lmKHQpe3ZhciByPUModGhpcykuY29uc3RydWN0b3I7bj1SZWZsZWN0LmNvbnN0cnVjdChvLGFyZ3VtZW50cyxyKX1lbHNlIG49by5hcHBseSh0aGlzLGFyZ3VtZW50cyk7cmV0dXJuIHYodGhpcyxuKX19ZnVuY3Rpb24gdihlLHQpe3JldHVybiF0fHwib2JqZWN0IiE9PWcodCkmJiJmdW5jdGlvbiIhPXR5cGVvZiB0P2Z1bmN0aW9uKGUpe2lmKHZvaWQgMD09PWUpdGhyb3cgbmV3IFJlZmVyZW5jZUVycm9yKCJ0aGlzIGhhc24ndCBiZWVuIGluaXRpYWxpc2VkIC0gc3VwZXIoKSBoYXNuJ3QgYmVlbiBjYWxsZWQiKTtyZXR1cm4gZX0oZSk6dH1mdW5jdGlvbiBDKGUpe3JldHVybihDPU9iamVjdC5zZXRQcm90b3R5cGVPZj9PYmplY3QuZ2V0UHJvdG90eXBlT2Y6ZnVuY3Rpb24oZSl7cmV0dXJuIGUuX19wcm90b19ffHxPYmplY3QuZ2V0UHJvdG90eXBlT2YoZSl9KShlKX1mdW5jdGlvbiB3KGUsdCl7dmFyIG49ImRhdGEtY2xpcGJvYXJkLSIuY29uY2F0KGUpO2lmKHQuaGFzQXR0cmlidXRlKG4pKXJldHVybiB0LmdldEF0dHJpYnV0ZShuKX12YXIgRT1mdW5jdGlvbihlKXshZnVuY3Rpb24oZSx0KXtpZigiZnVuY3Rpb24iIT10eXBlb2YgdCYmbnVsbCE9PXQpdGhyb3cgbmV3IFR5cGVFcnJvcigiU3VwZXIgZXhwcmVzc2lvbiBtdXN0IGVpdGhlciBiZSBudWxsIG9yIGEgZnVuY3Rpb24iKTtlLnByb3RvdHlwZT1PYmplY3QuY3JlYXRlKHQmJnQucHJvdG90eXBlLHtjb25zdHJ1Y3Rvcjp7dmFsdWU6ZSx3cml0YWJsZTohMCxjb25maWd1cmFibGU6ITB9fSksdCYmYihlLHQpfShuLGUpO3ZhciB0PW0obik7ZnVuY3Rpb24gbihlLG8pe3ZhciByO3JldHVybiBmdW5jdGlvbihlLHQpe2lmKCEoZSBpbnN0YW5jZW9mIHQpKXRocm93IG5ldyBUeXBlRXJyb3IoIkNhbm5vdCBjYWxsIGEgY2xhc3MgYXMgYSBmdW5jdGlvbiIpfSh0aGlzLG4pLChyPXQuY2FsbCh0aGlzKSkucmVzb2x2ZU9wdGlvbnMobyksci5saXN0ZW5DbGljayhlKSxyfXJldHVybiBmdW5jdGlvbihlLHQsbil7dCYmaChlLnByb3RvdHlwZSx0KSxuJiZoKGUsbil9KG4sW3trZXk6InJlc29sdmVPcHRpb25zIix2YWx1ZTpmdW5jdGlvbigpe3ZhciBlPWFyZ3VtZW50cy5sZW5ndGg+MCYmdm9pZCAwIT09YXJndW1lbnRzWzBdP2FyZ3VtZW50c1swXTp7fTt0aGlzLmFjdGlvbj0iZnVuY3Rpb24iPT10eXBlb2YgZS5hY3Rpb24/ZS5hY3Rpb246dGhpcy5kZWZhdWx0QWN0aW9uLHRoaXMudGFyZ2V0PSJmdW5jdGlvbiI9PXR5cGVvZiBlLnRhcmdldD9lLnRhcmdldDp0aGlzLmRlZmF1bHRUYXJnZXQsdGhpcy50ZXh0PSJmdW5jdGlvbiI9PXR5cGVvZiBlLnRleHQ/ZS50ZXh0OnRoaXMuZGVmYXVsdFRleHQsdGhpcy5jb250YWluZXI9Im9iamVjdCI9PT1nKGUuY29udGFpbmVyKT9lLmNvbnRhaW5lcjpkb2N1bWVudC5ib2R5fX0se2tleToibGlzdGVuQ2xpY2siLHZhbHVlOmZ1bmN0aW9uKGUpe3ZhciB0PXRoaXM7dGhpcy5saXN0ZW5lcj1pKCkoZSwiY2xpY2siLChmdW5jdGlvbihlKXtyZXR1cm4gdC5vbkNsaWNrKGUpfSkpfX0se2tleToib25DbGljayIsdmFsdWU6ZnVuY3Rpb24oZSl7dmFyIHQ9ZS5kZWxlZ2F0ZVRhcmdldHx8ZS5jdXJyZW50VGFyZ2V0LG49dGhpcy5hY3Rpb24odCl8fCJjb3B5IixvPXkoe2FjdGlvbjpuLGNvbnRhaW5lcjp0aGlzLmNvbnRhaW5lcix0YXJnZXQ6dGhpcy50YXJnZXQodCksdGV4dDp0aGlzLnRleHQodCl9KTt0aGlzLmVtaXQobz8ic3VjY2VzcyI6ImVycm9yIix7YWN0aW9uOm4sdGV4dDpvLHRyaWdnZXI6dCxjbGVhclNlbGVjdGlvbjpmdW5jdGlvbigpe3QmJnQuZm9jdXMoKSx3aW5kb3cuZ2V0U2VsZWN0aW9uKCkucmVtb3ZlQWxsUmFuZ2VzKCl9fSl9fSx7a2V5OiJkZWZhdWx0QWN0aW9uIix2YWx1ZTpmdW5jdGlvbihlKXtyZXR1cm4gdygiYWN0aW9uIixlKX19LHtrZXk6ImRlZmF1bHRUYXJnZXQiLHZhbHVlOmZ1bmN0aW9uKGUpe3ZhciB0PXcoInRhcmdldCIsZSk7aWYodClyZXR1cm4gZG9jdW1lbnQucXVlcnlTZWxlY3Rvcih0KX19LHtrZXk6ImRlZmF1bHRUZXh0Iix2YWx1ZTpmdW5jdGlvbihlKXtyZXR1cm4gdygidGV4dCIsZSl9fSx7a2V5OiJkZXN0cm95Iix2YWx1ZTpmdW5jdGlvbigpe3RoaXMubGlzdGVuZXIuZGVzdHJveSgpfX1dLFt7a2V5OiJjb3B5Iix2YWx1ZTpmdW5jdGlvbihlKXt2YXIgdD1hcmd1bWVudHMubGVuZ3RoPjEmJnZvaWQgMCE9PWFyZ3VtZW50c1sxXT9hcmd1bWVudHNbMV06e2NvbnRhaW5lcjpkb2N1bWVudC5ib2R5fTtyZXR1cm4gZChlLHQpfX0se2tleToiY3V0Iix2YWx1ZTpmdW5jdGlvbihlKXtyZXR1cm4gbChlKX19LHtrZXk6ImlzU3VwcG9ydGVkIix2YWx1ZTpmdW5jdGlvbigpe3ZhciBlPWFyZ3VtZW50cy5sZW5ndGg+MCYmdm9pZCAwIT09YXJndW1lbnRzWzBdP2FyZ3VtZW50c1swXTpbImNvcHkiLCJjdXQiXSx0PSJzdHJpbmciPT10eXBlb2YgZT9bZV06ZSxuPSEhZG9jdW1lbnQucXVlcnlDb21tYW5kU3VwcG9ydGVkO3JldHVybiB0LmZvckVhY2goKGZ1bmN0aW9uKGUpe249biYmISFkb2N1bWVudC5xdWVyeUNvbW1hbmRTdXBwb3J0ZWQoZSl9KSksbn19XSksbn0ocigpKX0sODI4OmZ1bmN0aW9uKGUpe2lmKCJ1bmRlZmluZWQiIT10eXBlb2YgRWxlbWVudCYmIUVsZW1lbnQucHJvdG90eXBlLm1hdGNoZXMpe3ZhciB0PUVsZW1lbnQucHJvdG90eXBlO3QubWF0Y2hlcz10Lm1hdGNoZXNTZWxlY3Rvcnx8dC5tb3pNYXRjaGVzU2VsZWN0b3J8fHQubXNNYXRjaGVzU2VsZWN0b3J8fHQub01hdGNoZXNTZWxlY3Rvcnx8dC53ZWJraXRNYXRjaGVzU2VsZWN0b3J9ZS5leHBvcnRzPWZ1bmN0aW9uKGUsdCl7Zm9yKDtlJiY5IT09ZS5ub2RlVHlwZTspe2lmKCJmdW5jdGlvbiI9PXR5cGVvZiBlLm1hdGNoZXMmJmUubWF0Y2hlcyh0KSlyZXR1cm4gZTtlPWUucGFyZW50Tm9kZX19fSw0Mzg6ZnVuY3Rpb24oZSx0LG4pe3ZhciBvPW4oODI4KTtmdW5jdGlvbiByKGUsdCxuLG8scil7dmFyIGk9cy5hcHBseSh0aGlzLGFyZ3VtZW50cyk7cmV0dXJuIGUuYWRkRXZlbnRMaXN0ZW5lcihuLGkscikse2Rlc3Ryb3k6ZnVuY3Rpb24oKXtlLnJlbW92ZUV2ZW50TGlzdGVuZXIobixpLHIpfX19ZnVuY3Rpb24gcyhlLHQsbixyKXtyZXR1cm4gZnVuY3Rpb24obil7bi5kZWxlZ2F0ZVRhcmdldD1vKG4udGFyZ2V0LHQpLG4uZGVsZWdhdGVUYXJnZXQmJnIuY2FsbChlLG4pfX1lLmV4cG9ydHM9ZnVuY3Rpb24oZSx0LG4sbyxzKXtyZXR1cm4iZnVuY3Rpb24iPT10eXBlb2YgZS5hZGRFdmVudExpc3RlbmVyP3IuYXBwbHkobnVsbCxhcmd1bWVudHMpOiJmdW5jdGlvbiI9PXR5cGVvZiBuP3IuYmluZChudWxsLGRvY3VtZW50KS5hcHBseShudWxsLGFyZ3VtZW50cyk6KCJzdHJpbmciPT10eXBlb2YgZSYmKGU9ZG9jdW1lbnQucXVlcnlTZWxlY3RvckFsbChlKSksQXJyYXkucHJvdG90eXBlLm1hcC5jYWxsKGUsKGZ1bmN0aW9uKGUpe3JldHVybiByKGUsdCxuLG8scyl9KSkpfX0sODc5OmZ1bmN0aW9uKGUsdCl7dC5ub2RlPWZ1bmN0aW9uKGUpe3JldHVybiB2b2lkIDAhPT1lJiZlIGluc3RhbmNlb2YgSFRNTEVsZW1lbnQmJjE9PT1lLm5vZGVUeXBlfSx0Lm5vZGVMaXN0PWZ1bmN0aW9uKGUpe3ZhciBuPU9iamVjdC5wcm90b3R5cGUudG9TdHJpbmcuY2FsbChlKTtyZXR1cm4gdm9pZCAwIT09ZSYmKCJbb2JqZWN0IE5vZGVMaXN0XSI9PT1ufHwiW29iamVjdCBIVE1MQ29sbGVjdGlvbl0iPT09bikmJiJsZW5ndGgiaW4gZSYmKDA9PT1lLmxlbmd0aHx8dC5ub2RlKGVbMF0pKX0sdC5zdHJpbmc9ZnVuY3Rpb24oZSl7cmV0dXJuInN0cmluZyI9PXR5cGVvZiBlfHxlIGluc3RhbmNlb2YgU3RyaW5nfSx0LmZuPWZ1bmN0aW9uKGUpe3JldHVybiJbb2JqZWN0IEZ1bmN0aW9uXSI9PT1PYmplY3QucHJvdG90eXBlLnRvU3RyaW5nLmNhbGwoZSl9fSwzNzA6ZnVuY3Rpb24oZSx0LG4pe3ZhciBvPW4oODc5KSxyPW4oNDM4KTtlLmV4cG9ydHM9ZnVuY3Rpb24oZSx0LG4pe2lmKCFlJiYhdCYmIW4pdGhyb3cgbmV3IEVycm9yKCJNaXNzaW5nIHJlcXVpcmVkIGFyZ3VtZW50cyIpO2lmKCFvLnN0cmluZyh0KSl0aHJvdyBuZXcgVHlwZUVycm9yKCJTZWNvbmQgYXJndW1lbnQgbXVzdCBiZSBhIFN0cmluZyIpO2lmKCFvLmZuKG4pKXRocm93IG5ldyBUeXBlRXJyb3IoIlRoaXJkIGFyZ3VtZW50IG11c3QgYmUgYSBGdW5jdGlvbiIpO2lmKG8ubm9kZShlKSlyZXR1cm4gZnVuY3Rpb24oZSx0LG4pe3JldHVybiBlLmFkZEV2ZW50TGlzdGVuZXIodCxuKSx7ZGVzdHJveTpmdW5jdGlvbigpe2UucmVtb3ZlRXZlbnRMaXN0ZW5lcih0LG4pfX19KGUsdCxuKTtpZihvLm5vZGVMaXN0KGUpKXJldHVybiBmdW5jdGlvbihlLHQsbil7cmV0dXJuIEFycmF5LnByb3RvdHlwZS5mb3JFYWNoLmNhbGwoZSwoZnVuY3Rpb24oZSl7ZS5hZGRFdmVudExpc3RlbmVyKHQsbil9KSkse2Rlc3Ryb3k6ZnVuY3Rpb24oKXtBcnJheS5wcm90b3R5cGUuZm9yRWFjaC5jYWxsKGUsKGZ1bmN0aW9uKGUpe2UucmVtb3ZlRXZlbnRMaXN0ZW5lcih0LG4pfSkpfX19KGUsdCxuKTtpZihvLnN0cmluZyhlKSlyZXR1cm4gZnVuY3Rpb24oZSx0LG4pe3JldHVybiByKGRvY3VtZW50LmJvZHksZSx0LG4pfShlLHQsbik7dGhyb3cgbmV3IFR5cGVFcnJvcigiRmlyc3QgYXJndW1lbnQgbXVzdCBiZSBhIFN0cmluZywgSFRNTEVsZW1lbnQsIEhUTUxDb2xsZWN0aW9uLCBvciBOb2RlTGlzdCIpfX0sODE3OmZ1bmN0aW9uKGUpe2UuZXhwb3J0cz1mdW5jdGlvbihlKXt2YXIgdDtpZigiU0VMRUNUIj09PWUubm9kZU5hbWUpZS5mb2N1cygpLHQ9ZS52YWx1ZTtlbHNlIGlmKCJJTlBVVCI9PT1lLm5vZGVOYW1lfHwiVEVYVEFSRUEiPT09ZS5ub2RlTmFtZSl7dmFyIG49ZS5oYXNBdHRyaWJ1dGUoInJlYWRvbmx5Iik7bnx8ZS5zZXRBdHRyaWJ1dGUoInJlYWRvbmx5IiwiIiksZS5zZWxlY3QoKSxlLnNldFNlbGVjdGlvblJhbmdlKDAsZS52YWx1ZS5sZW5ndGgpLG58fGUucmVtb3ZlQXR0cmlidXRlKCJyZWFkb25seSIpLHQ9ZS52YWx1ZX1lbHNle2UuaGFzQXR0cmlidXRlKCJjb250ZW50ZWRpdGFibGUiKSYmZS5mb2N1cygpO3ZhciBvPXdpbmRvdy5nZXRTZWxlY3Rpb24oKSxyPWRvY3VtZW50LmNyZWF0ZVJhbmdlKCk7ci5zZWxlY3ROb2RlQ29udGVudHMoZSksby5yZW1vdmVBbGxSYW5nZXMoKSxvLmFkZFJhbmdlKHIpLHQ9by50b1N0cmluZygpfXJldHVybiB0fX0sMjc5OmZ1bmN0aW9uKGUpe2Z1bmN0aW9uIHQoKXt9dC5wcm90b3R5cGU9e29uOmZ1bmN0aW9uKGUsdCxuKXt2YXIgbz10aGlzLmV8fCh0aGlzLmU9e30pO3JldHVybihvW2VdfHwob1tlXT1bXSkpLnB1c2goe2ZuOnQsY3R4Om59KSx0aGlzfSxvbmNlOmZ1bmN0aW9uKGUsdCxuKXt2YXIgbz10aGlzO2Z1bmN0aW9uIHIoKXtvLm9mZihlLHIpLHQuYXBwbHkobixhcmd1bWVudHMpfXJldHVybiByLl89dCx0aGlzLm9uKGUscixuKX0sZW1pdDpmdW5jdGlvbihlKXtmb3IodmFyIHQ9W10uc2xpY2UuY2FsbChhcmd1bWVudHMsMSksbj0oKHRoaXMuZXx8KHRoaXMuZT17fSkpW2VdfHxbXSkuc2xpY2UoKSxvPTAscj1uLmxlbmd0aDtvPHI7bysrKW5bb10uZm4uYXBwbHkobltvXS5jdHgsdCk7cmV0dXJuIHRoaXN9LG9mZjpmdW5jdGlvbihlLHQpe3ZhciBuPXRoaXMuZXx8KHRoaXMuZT17fSksbz1uW2VdLHI9W107aWYobyYmdClmb3IodmFyIHM9MCxpPW8ubGVuZ3RoO3M8aTtzKyspb1tzXS5mbiE9PXQmJm9bc10uZm4uXyE9PXQmJnIucHVzaChvW3NdKTtyZXR1cm4gci5sZW5ndGg/bltlXT1yOmRlbGV0ZSBuW2VdLHRoaXN9fSxlLmV4cG9ydHM9dCxlLmV4cG9ydHMuVGlueUVtaXR0ZXI9dH19LHQ9e307ZnVuY3Rpb24gbihvKXtpZih0W29dKXJldHVybiB0W29dLmV4cG9ydHM7dmFyIHI9dFtvXT17ZXhwb3J0czp7fX07cmV0dXJuIGVbb10ocixyLmV4cG9ydHMsbiksci5leHBvcnRzfXJldHVybiBuLm49ZnVuY3Rpb24oZSl7dmFyIHQ9ZSYmZS5fX2VzTW9kdWxlP2Z1bmN0aW9uKCl7cmV0dXJuIGUuZGVmYXVsdH06ZnVuY3Rpb24oKXtyZXR1cm4gZX07cmV0dXJuIG4uZCh0LHthOnR9KSx0fSxuLmQ9ZnVuY3Rpb24oZSx0KXtmb3IodmFyIG8gaW4gdCluLm8odCxvKSYmIW4ubyhlLG8pJiZPYmplY3QuZGVmaW5lUHJvcGVydHkoZSxvLHtlbnVtZXJhYmxlOiEwLGdldDp0W29dfSl9LG4ubz1mdW5jdGlvbihlLHQpe3JldHVybiBPYmplY3QucHJvdG90eXBlLmhhc093blByb3BlcnR5LmNhbGwoZSx0KX0sbig2ODYpfSgpLmRlZmF1bHR9fV0pOw==")))
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
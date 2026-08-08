%global debug_package %{nil}

Name: protonbox
Version: 0.1.0
Release: 1%{?dist}
Summary: Useful Proton launch commands manager
License: MIT
URL: https://github.com/LucianoSkx/protonbox
Source0: %{name}-%{version}.tar.gz
BuildArch: x86_64

%description
ProtonBox is a simple GUI to browse, copy and combine useful Proton launch
commands (standard Proton, Proton-GE and Proton-CachyOS) for Steam.

%prep
%setup -q

%install
mkdir -p %{buildroot}%{_bindir}
install -m755 protonbox %{buildroot}%{_bindir}/protonbox
mkdir -p %{buildroot}%{_datadir}/applications
install -m644 ProtonBox.desktop %{buildroot}%{_datadir}/applications/ProtonBox.desktop
mkdir -p %{buildroot}%{_datadir}/icons/hicolor/256x256/apps
install -m644 icon.png %{buildroot}%{_datadir}/icons/hicolor/256x256/apps/protonbox.png

%files
%{_bindir}/protonbox
%{_datadir}/applications/ProtonBox.desktop
%{_datadir}/icons/hicolor/256x256/apps/protonbox.png

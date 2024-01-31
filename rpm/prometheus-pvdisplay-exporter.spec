%define        __spec_install_post %{nil}
%define          debug_package %{nil}
%define        __os_install_post %{_dbpath}/brp-compress

Name:           prometheus-pvdisplay-exporter
Version:        VERSION
Release:        1%{?dist}
Summary:        Prometheus exporter for pvdisplay
Group:          Monitoring

License:        GPL 3.0
URL:            https://github.com/GSI-HPC/prometheus-pvdisplay-exporter
Source0:        %{name}-%{version}.tar.gz

Requires(post): systemd
Requires(preun): systemd
Requires(postun): systemd
%{?systemd_requires}
BuildRequires:  systemd

BuildRoot:      %{_tmppath}/%{name}-%{version}-1-root

%description
Prometheus exporter for pvdisplay

%prep
%setup -q

%build
# Empty section.

%install
rm -rf %{buildroot}
mkdir -p %{buildroot}%{_unitdir}/
cp usr/lib/systemd/system/%{name}.service %{buildroot}%{_unitdir}/

# in builddir
cp -a * %{buildroot}

%clean
rm -rf %{buildroot}

%post
systemctl enable %{name}.service
systemctl start %{name}.service

%preun
%systemd_preun %{name}.service

%postun
%systemd_postun_with_restart %{name}.service

%files
%defattr(-,root,root,-)
%{_sbindir}/prometheus-pvdisplay-exporter
%{_unitdir}/%{name}.service
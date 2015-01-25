#FROM debian:jessie
FROM jprjr/arch:latest
MAINTAINER William McGann <contact@WilliamMcGann.com>

COPY root/etc/pacman.conf /etc/pacman.conf

RUN pacman -Syu --noconfirm \
	&& pacman-db-upgrade \
	&& pacman -S --noconfirm --quiet --needed base-devel curl musl

RUN mkdir /{build,output}

COPY root/ /

CMD ["build-s6"]
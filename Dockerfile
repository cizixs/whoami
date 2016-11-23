FROM scratch
ADD whoami /whoami
EXPOSE 3000
CMD ["/whoami"]

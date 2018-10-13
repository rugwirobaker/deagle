FROM scratch

COPY bin/* ./

EXPOSE 5000

ENTRYPOINT ["./deagle"]




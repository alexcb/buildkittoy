# buildkittoy
playing around with buildkit



    alex@mah:~/gh/alexcb/buildkittoy$ docker rmi alextest; go run cmd/toyapp/main.go
    Error: No such image: alextest
    running!
    Enter text: blahhhhhh
    status is &{[0xc000118a80 0xc000118ae0 0xc000118b40] [] []}
     vertex: &{sha256:1b8e7bb36a9335a1ba979e3d548d03af76a0a293cdd7b99e1269f799bbeefad0 [] docker-image://docker.io/library/golang:1.13-alpine <nil> <nil> false }
     vertex: &{sha256:6593ce8dfe9da193bbe6fb8525f0b27f4a0b05a566f098422b09d55964693f32 [sha256:1b8e7bb36a9335a1ba979e3d548d03af76a0a293cdd7b99e1269f799bbeefad0] mkfile /file <nil> <nil> false }
     vertex: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec [sha256:6593ce8dfe9da193bbe6fb8525f0b27f4a0b05a566f098422b09d55964693f32] apk add --no-cache git <nil> <nil> false }
    status is &{[0xc000118ba0] [0xc000126460] []}
     vertex: &{sha256:1b8e7bb36a9335a1ba979e3d548d03af76a0a293cdd7b99e1269f799bbeefad0 [] docker-image://docker.io/library/golang:1.13-alpine 2020-07-03 21:21:16.974967723 +0000 UTC <nil> false }
     status: &{resolve docker.io/library/golang:1.13-alpine sha256:1b8e7bb36a9335a1ba979e3d548d03af76a0a293cdd7b99e1269f799bbeefad0  0 0 2020-07-03 21:21:16.975033264 +0000 UTC 2020-07-03 21:21:16.975032859 +0000 UTC <nil>}
    status is &{[0xc000118fc0] [0xc0001267e0] []}
     vertex: &{sha256:1b8e7bb36a9335a1ba979e3d548d03af76a0a293cdd7b99e1269f799bbeefad0 [] docker-image://docker.io/library/golang:1.13-alpine 2020-07-03 21:21:16.974967723 +0000 UTC 2020-07-03 21:21:18.091186285 +0000 UTC false }
     status: &{resolve docker.io/library/golang:1.13-alpine sha256:1b8e7bb36a9335a1ba979e3d548d03af76a0a293cdd7b99e1269f799bbeefad0  0 0 2020-07-03 21:21:18.091103621 +0000 UTC 2020-07-03 21:21:16.975032859 +0000 UTC 2020-07-03 21:21:18.091101831 +0000 UTC}
    status is &{[0xc000119080] [] []}
     vertex: &{sha256:1b8e7bb36a9335a1ba979e3d548d03af76a0a293cdd7b99e1269f799bbeefad0 [] docker-image://docker.io/library/golang:1.13-alpine 2020-07-03 21:21:18.091751777 +0000 UTC <nil> true }
    status is &{[0xc0001190e0] [] []}
     vertex: &{sha256:1b8e7bb36a9335a1ba979e3d548d03af76a0a293cdd7b99e1269f799bbeefad0 [] docker-image://docker.io/library/golang:1.13-alpine 2020-07-03 21:21:18.091751777 +0000 UTC 2020-07-03 21:21:18.091894768 +0000 UTC true }
    status is &{[0xc000119200] [] []}
     vertex: &{sha256:6593ce8dfe9da193bbe6fb8525f0b27f4a0b05a566f098422b09d55964693f32 [sha256:1b8e7bb36a9335a1ba979e3d548d03af76a0a293cdd7b99e1269f799bbeefad0] mkfile /file 2020-07-03 21:21:18.092182593 +0000 UTC <nil> false }
    status is &{[0xc0005a21e0] [] []}
     vertex: &{sha256:6593ce8dfe9da193bbe6fb8525f0b27f4a0b05a566f098422b09d55964693f32 [sha256:1b8e7bb36a9335a1ba979e3d548d03af76a0a293cdd7b99e1269f799bbeefad0] mkfile /file 2020-07-03 21:21:18.092182593 +0000 UTC 2020-07-03 21:21:18.21404892 +0000 UTC false }
    status is &{[0xc0005a22a0] [] []}
     vertex: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec [sha256:6593ce8dfe9da193bbe6fb8525f0b27f4a0b05a566f098422b09d55964693f32] apk add --no-cache git 2020-07-03 21:21:18.218201066 +0000 UTC <nil> false }
    status is &{[] [] [0xc0005a65a0]}
     log: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec 1 [102 101 116 99 104 32 104 116 116 112 58 47 47 100 108 45 99 100 110 46 97 108 112 105 110 101 108 105 110 117 120 46 111 114 103 47 97 108 112 105 110 101 47 118 51 46 49 50 47 109 97 105 110 47 120 56 54 95 54 52 47 65 80 75 73 78 68 69 88 46 116 97 114 46 103 122 10] 2020-07-03 21:21:18.420384041 +0000 UTC}
    status is &{[] [] [0xc0005a66e0]}
     log: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec 1 [102 101 116 99 104 32 104 116 116 112 58 47 47 100 108 45 99 100 110 46 97 108 112 105 110 101 108 105 110 117 120 46 111 114 103 47 97 108 112 105 110 101 47 118 51 46 49 50 47 99 111 109 109 117 110 105 116 121 47 120 56 54 95 54 52 47 65 80 75 73 78 68 69 88 46 116 97 114 46 103 122 10] 2020-07-03 21:21:18.796040925 +0000 UTC}
    status is &{[] [] [0xc00048c190]}
     log: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec 1 [40 49 47 53 41 32 73 110 115 116 97 108 108 105 110 103 32 110 103 104 116 116 112 50 45 108 105 98 115 32 40 49 46 52 49 46 48 45 114 48 41 10] 2020-07-03 21:21:19.239350504 +0000 UTC}
    status is &{[] [] [0xc00048c2d0]}
     log: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec 1 [40 50 47 53 41 32 73 110 115 116 97 108 108 105 110 103 32 108 105 98 99 117 114 108 32 40 55 46 54 57 46 49 45 114 48 41 10] 2020-07-03 21:21:19.291372707 +0000 UTC}
    status is &{[] [] [0xc0000900f0]}
     log: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec 1 [40 51 47 53 41 32 73 110 115 116 97 108 108 105 110 103 32 101 120 112 97 116 32 40 50 46 50 46 57 45 114 49 41 10] 2020-07-03 21:21:19.388078623 +0000 UTC}
    status is &{[] [] [0xc00048c460]}
     log: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec 1 [40 52 47 53 41 32 73 110 115 116 97 108 108 105 110 103 32 112 99 114 101 50 32 40 49 48 46 51 53 45 114 48 41 10] 2020-07-03 21:21:19.42691752 +0000 UTC}
    status is &{[] [] [0xc00048c5a0]}
     log: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec 1 [40 53 47 53 41 32 73 110 115 116 97 108 108 105 110 103 32 103 105 116 32 40 50 46 50 54 46 50 45 114 48 41 10] 2020-07-03 21:21:19.54342778 +0000 UTC}
    status is &{[] [] [0xc00012a0a0]}
     log: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec 1 [69 120 101 99 117 116 105 110 103 32 98 117 115 121 98 111 120 45 49 46 51 49 46 49 45 114 49 54 46 116 114 105 103 103 101 114 10] 2020-07-03 21:21:21.470106981 +0000 UTC}
    status is &{[] [] [0xc00012a280]}
     log: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec 1 [79 75 58 32 50 50 32 77 105 66 32 105 110 32 50 48 32 112 97 99 107 97 103 101 115 10] 2020-07-03 21:21:21.475806505 +0000 UTC}
    status is &{[0xc00007e360] [] []}
     vertex: &{sha256:c58c318dd3741227d9a0de968257e163c8dc09fce40ac09ed82bcddf9ccf47ec [sha256:6593ce8dfe9da193bbe6fb8525f0b27f4a0b05a566f098422b09d55964693f32] apk add --no-cache git 2020-07-03 21:21:18.218201066 +0000 UTC 2020-07-03 21:21:21.582760565 +0000 UTC false }
    status is &{[0xc00007e420] [] []}
     vertex: &{sha256:69a2560eef4d3ece902c3b5149d142e9bd132f25db0bc9e35b94201534c415d2 [] exporting to oci image format 2020-07-03 21:21:21.584544498 +0000 UTC <nil> false }
    status is &{[] [0xc0001aa2a0] []}
     status: &{exporting layers sha256:69a2560eef4d3ece902c3b5149d142e9bd132f25db0bc9e35b94201534c415d2  0 0 2020-07-03 21:21:21.585683486 +0000 UTC 2020-07-03 21:21:21.585683053 +0000 UTC <nil>}
    status is &{[] [0xc0000c4310] []}
     status: &{exporting layers sha256:69a2560eef4d3ece902c3b5149d142e9bd132f25db0bc9e35b94201534c415d2  0 0 2020-07-03 21:21:23.350203274 +0000 UTC 2020-07-03 21:21:21.585683053 +0000 UTC 2020-07-03 21:21:23.350202764 +0000 UTC}
    status is &{[] [0xc0000c4460] []}
     status: &{exporting manifest sha256:05d12cad440c4d8765a56cdf63e4430757754e75513fa78f78366bd174a4a544 sha256:69a2560eef4d3ece902c3b5149d142e9bd132f25db0bc9e35b94201534c415d2  0 0 2020-07-03 21:21:23.354038698 +0000 UTC 2020-07-03 21:21:23.354038348 +0000 UTC <nil>}
    status is &{[] [0xc0001a41c0 0xc0001a4230] []}
     status: &{exporting manifest sha256:05d12cad440c4d8765a56cdf63e4430757754e75513fa78f78366bd174a4a544 sha256:69a2560eef4d3ece902c3b5149d142e9bd132f25db0bc9e35b94201534c415d2  0 0 2020-07-03 21:21:23.36215484 +0000 UTC 2020-07-03 21:21:23.354038348 +0000 UTC 2020-07-03 21:21:23.362154241 +0000 UTC}
     status: &{exporting config sha256:431d46178b7dbdeadf145470473d7c625201d87c479a23599749582388eb74bc sha256:69a2560eef4d3ece902c3b5149d142e9bd132f25db0bc9e35b94201534c415d2  0 0 2020-07-03 21:21:23.362166015 +0000 UTC 2020-07-03 21:21:23.362165737 +0000 UTC <nil>}
    status is &{[] [0xc0002fe5b0] []}
     status: &{exporting config sha256:431d46178b7dbdeadf145470473d7c625201d87c479a23599749582388eb74bc sha256:69a2560eef4d3ece902c3b5149d142e9bd132f25db0bc9e35b94201534c415d2  0 0 2020-07-03 21:21:23.369129347 +0000 UTC 2020-07-03 21:21:23.362165737 +0000 UTC 2020-07-03 21:21:23.369127824 +0000 UTC}
    status is &{[] [0xc0002fe770] []}
     status: &{sending tarball sha256:69a2560eef4d3ece902c3b5149d142e9bd132f25db0bc9e35b94201534c415d2  0 0 2020-07-03 21:21:23.369275113 +0000 UTC 2020-07-03 21:21:23.369274849 +0000 UTC <nil>}
    status is &{[] [0xc0000c42a0] []}
     status: &{sending tarball sha256:69a2560eef4d3ece902c3b5149d142e9bd132f25db0bc9e35b94201534c415d2  0 0 2020-07-03 21:21:24.056479797 +0000 UTC 2020-07-03 21:21:23.369274849 +0000 UTC 2020-07-03 21:21:24.056479233 +0000 UTC}
    4f2a9fe9a6fe: Loading layer [==================================================>]     899B/899B
    status is &{[0xc0002a6300] [] []}
     vertex: &{sha256:69a2560eef4d3ece902c3b5149d142e9bd132f25db0bc9e35b94201534c415d2 [] exporting to oci image format 2020-07-03 21:21:21.584544498 +0000 UTC 2020-07-03 21:21:24.059791347 +0000 UTC false }
    fb770f36372c: Loading layer [==================================================>]  8.311MB/8.311MB
    Loaded image: alextest:latest
    done loading into docker
    finished!


Checking out the contents of the file

    alex@mah:~/gh/alexcb/buildkittoy$ docker run --rm alextest:latest /bin/cat /file
    blahhhhhh

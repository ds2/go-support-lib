plugins {
    id("net.researchgate.release") version "2.8.1"
    id("com.pivotstir.gogradle") version "1.1.5"
}

extensions.getByType(com.pivotstir.gogradle.GoPluginExtension::class).apply {
    env {
        version = "1.16"
        useSandbox = false
    }

    /*
    build {
        packagePaths = listOf("./...")
    }
    */

    dep {
        thirdpartyIgnored = true
    }

    dependencies {
        build("github.com/golang/protobuf@v1.4.2")
        test("github.com/stretchr/testify@v1.5.1")
    }
}

tasks.register<com.pivotstir.gogradle.tasks.GoBuild>("build") {
    group = "Build"
    description = "Dummy mirror for goBuild"
}

tasks.register<com.pivotstir.gogradle.tasks.GoClean>("clean") {
    group = "Build"
    description = "Dummy mirror for goClean"
}

release {
    tagTemplate = "v\${version}"
    with(propertyMissing("git") as net.researchgate.release.GitAdapter.GitConfig) {
        requireBranch = "releases/2021"
    }
}

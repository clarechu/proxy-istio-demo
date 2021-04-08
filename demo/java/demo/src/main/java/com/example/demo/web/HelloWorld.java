package com.example.demo.web;

import com.google.common.io.CharStreams;
import org.springframework.context.ResourceLoaderAware;
import org.springframework.core.io.ClassPathResource;
import org.springframework.core.io.Resource;
import org.springframework.core.io.ResourceLoader;
import org.springframework.util.ResourceUtils;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.beans.factory.annotation.Autowired;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.*;
import java.util.HashMap;
import java.util.Map;

@RestController
@RequestMapping(value = "/")
public class HelloWorld implements ResourceLoaderAware {

    @GetMapping(value = "/")
    public Object Hello(HttpServletRequest request, HttpServletResponse response) {
        String proxyName = request.getHeader("proxy");
        System.out.printf("proxy name: %s \n", proxyName);
        response.addHeader("demo ->", "java");
        Map<String, Object> resp = new HashMap<>();
        resp.put("up", true);
        return resp;
    }
    private ResourceLoader resourceLoader;

    @GetMapping(value = "/test")
    public Object Test() throws IOException {
        Resource resource = resourceLoader.getResource("classpath:static/test.yaml");
        InputStream inputStream = resource.getInputStream(); // <-- this is the difference
        String text = null;
        try (Reader reader = new InputStreamReader(inputStream)) {
            text = CharStreams.toString(reader);
        }
        return text;
    }

    @Override
    public void setResourceLoader(ResourceLoader resourceLoader) {
        this.resourceLoader = resourceLoader;
    }
}
package com.example.demo.web;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.HashMap;
import java.util.Map;

@RestController
@RequestMapping(value = "/")
public class HelloWorld {

    @GetMapping(value = "/")
    public Object Hello(HttpServletRequest request, HttpServletResponse response) {
        String proxyName = request.getHeader("proxy");
        System.out.printf("proxy name: %s \n", proxyName);
        response.addHeader("demo ->", "java");
        Map<String, Object> resp = new HashMap<>();
        resp.put("up", true);
        return resp;
    }


}

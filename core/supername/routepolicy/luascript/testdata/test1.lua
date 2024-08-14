if MatchFuncs == nil then
	MatchFuncs = {}
end

MatchFuncs["www.test1.com"] = function(ctx, clusters)
	DefaultLabelsMatch = function(ctx, labels)
		if ctx["X-Zone"] ~= labels["X-Zone"] then
			return false
		end
		if labels["X-Lane"] ~= "default" then
			return false
		end
		return true
	end

	for _, cluster in pairs(clusters) do
		if DefaultLabelsMatch(ctx, cluster.Labels) then
			return { {cluster.Name, 1} }
		end
	end

	return {}
end
